package ddz

import (
	"fmt"
	"sort"
)

type CARD_TYPE uint32

const(
	ROCKET_TYPE 		CARD_TYPE= 1 	//火箭类型
	ZHADAN_TYPE 		CARD_TYPE= 2 	//炸弹类型
	DOULEL_INK_TYPE 	CARD_TYPE= 3 	//双顺类型
	SINGLE_LINK_TYPE 	CARD_TYPE= 4 	//单顺类型
	THREE_ONE_TYPE 		CARD_TYPE= 5 	//三带一类型
	THREE_TYPE 			CARD_TYPE= 6 	//三牌类型
	DOUBLE_TYPE 		CARD_TYPE= 7 	//对牌类型
	SINGLE_TYPE 		CARD_TYPE= 8 	//单牌类型
)

const(
	SMALL_GHOST = 80
	BIG_GHOST   = 100
)

//三牌
type ThreeCarder struct{
	//eg 333
	ThreeCard  int // 3
}

//三带一
type ThreeOneCarder struct{
	//eg 333带4
	ThreeCard  int // 3
	SingleCard int // 4
}

//三带二
type ThreeTwoCarder struct{
	//eg 333带44
	ThreeCard  int // 3
	DoubleCard int // 4
}


//单顺组
type SingleLinker struct{
	//eg 4 5 6 7 8 9 10 11
	BeginValue int //开始值  4
	Len        int //单顺长度 8
	Data      []int//牌内容 4 5 6 7 8 9 10 11
}

//双顺组
type DoubleLinker struct{
	//eg 44 55 66
	BeginValue int //4
	Len 	   int //6
	Data       []int //44 55 66
}

type Parser struct{
	//原始牌
	Cards []int

	M1 []int
	M2 []int
	M3 []int
	M4 []int

	SingleCards     []int 	//单排组
	DoubleCards     []int	//对牌组
	//注意:三带一 和 三带二 同时只能有一种,从代码看 遍历组装三带一时,最后M3数组被清空了(根据情况选用三带一 和 三带二)
	ThreeOneCards   []ThreeOneCarder    //三带一
	ThreeTwoCards   []ThreeTwoCarder    //三带二
	ThreeCards      []ThreeCarder       //三牌
	SingleLink	 	[]SingleLinker	//单顺组
	DoubleLink   	[]DoubleLinker   //双顺组
	Bombes       	[]int         //炸弹组 存放的是 4  10 12 等等
	Rocket       	[]int        //火箭组 存放的肯定是 14 小鬼 和 15 大鬼 如果只有其中一个,就存放到单牌组

}

/*
拆分牌型组顺序
火箭组(14 15) 炸弹组(4444) 双顺组(44 55 66)  单顺组(34567)  三带一组(3334) 三牌组(333) 对牌组(44) 单牌组(4 6 7 9)
*/

func (self* Parser) SplitCard(){
	//通过拆牌顺序 可以调整出不同的出牌策略
	self.fillRocket()
	self.fillZhaDan()
	self.fillDoubleLink()
	self.fillSingleLink()
	self.fillThreeCard()
	self.fillDoubleCard()
	self.fillSingleCard()
}

//组装单牌 3 5 11 12 (通过拆火箭 炸弹 双顺 单顺 三牌 对牌后 M1中剩下的就都是单牌了)
func (self* Parser) fillSingleCard(){

	self.SingleCards = append(self.SingleCards,self.M1...)
}

//组装对子牌 55 99 1111
func (self* Parser) fillDoubleCard() {

	for _,m2_v := range self.M2{

		self.DoubleCards = append(self.DoubleCards,m2_v)

		del_index := -1
		for m1_i,m1_v := range self.M1{
			if m2_v == m1_v{
				del_index = m1_i
				break
			}
		}
		if del_index != -1{
			self.M1 = append(self.M1[0:del_index], self.M1[del_index+1:]...)
		}
	}

	//清空M2
	self.M2 = self.M2[0:0]

}

//组装火箭牌 从单牌M1 中检测是否有火箭(即是大小鬼 14 15)
func (self* Parser) fillRocket(){
	var it,jt int
	var  have_small_gui,have_big_gui bool
	for i,v := range self.M1{
		if v == SMALL_GHOST{
			it = i
			have_small_gui= true
		}else if v == BIG_GHOST{
			jt = i
			have_big_gui = true
		}
	}

	if have_small_gui && have_big_gui{
		//大小鬼填充到火箭组
		self.Rocket = append(self.Rocket,self.M1[it],self.M1[jt])
		//remove M1 中的大小鬼
		self.M1 = append(self.M1[0:it],self.M1[it+1:]...)
		self.M1 = append(self.M1[0:jt],self.M1[jt+1:]...)
	}
}

//组装炸弹牌 5555 9999
func (self* Parser) fillZhaDan(){
	for i := 0; i < len(self.M4); i++{

		zhadan := self.M4[i]

		self.Bombes = append(self.Bombes,zhadan)

		var index int
		//remove M1 M2 M3 M4 中的炸弹牌
		//移除self.M1中的炸弹牌 begin
		for i,value := range self.M1{
			if value == zhadan{
				index = i
				break
			}
		}
		self.M1 = append(self.M1[0:index],self.M1[index+1:]...)

		//移除self.M2中的炸弹牌 begin
		for i,value := range self.M2{
			if value == zhadan{
				index = i
				break
			}
		}
		self.M2 = append(self.M2[0:index],self.M2[index+1:]...)

		//移除self.M3中的炸弹牌 begin
		for i,value := range self.M3{
			if value == zhadan{
				index = i
				break
			}
		}
		self.M3 = append(self.M3[0:index],self.M3[index+1:]...)
		//end
	}

	//M4已经拆完牌 清空
	self.M4 = self.M4[0:0]
}

func (self* Parser) fillDoubleLink()  {
	//组装双顺牌(分析M2中是否有超过3张连续的牌 eg:112233 66778899)
	for{
		sub_series_arr := MaxSerialArr(self.M2)
		length := len(sub_series_arr)
		if length >= 3{
			//生成一个双顺牌存放起来,同时移除M1 M2中对应数字,继续找下个双顺牌,找不到就退出
			node := &DoubleLinker{
				BeginValue:sub_series_arr[0],
				Len:length,
				Data:sub_series_arr,
			}
			self.DoubleLink = append(self.DoubleLink,*node)

			for _,v :=range sub_series_arr{

				for m1_i,m1_v := range self.M1{
					if m1_v == v{
						self.M1 = append(self.M1[0:m1_i],self.M1[m1_i+1:]...)
						break
					}
				}

				for m2_i,m2_v := range self.M2{
					if m2_v == v{
						self.M2 = append(self.M2[0:m2_i],self.M2[m2_i+1:]...)
						break
					}
				}
			}

			//继续找双顺牌
			continue
		}else{
			break
		}
	}
}

func (self* Parser) fillSingleLink()  {
	//组装单顺牌(分析M1中是否有5张及以上连续的牌 eg:1 2 3 4 5 or 8 9 10 11 12 13)
	for{
		sub_series_arr := MaxSerialArr(self.M1)
		length := len(sub_series_arr)
		if length >= 5{
			//生成一个单顺牌存放起来,同时移除M1中对应数字,继续找下个双顺牌,找不到就退出
			node := &SingleLinker{
				BeginValue:sub_series_arr[0],
				Len:length,
				Data:sub_series_arr,
			}
			self.SingleLink = append(self.SingleLink,*node)

			for _,v :=range sub_series_arr{

				for m1_i,m1_v := range self.M1{
					if m1_v == v{
						self.M1 = append(self.M1[0:m1_i],self.M1[m1_i+1:]...)
						break
					}
				}
			}

			//继续找单顺牌
			continue
		}else{
			break
		}
	}
}

func (self* Parser) fillThreeCard(){
	//组装三带一牌(从M3开始循环组装三带一牌,同时移除M3 M2 M1中都存在的牌以及移除M1中带的单牌)
	//要确保数组寻址有效(不然程序崩溃)
	for{
		for _,m3_v := range self.M3{
			//M1不等于0 就可以组成三带一
			if len(self.M1) != 0{
				node := &ThreeOneCarder{
					ThreeCard:m3_v,
					SingleCard:self.M1[0],//带牌用单牌组的头一个
				}
				self.ThreeOneCards = append(self.ThreeOneCards,*node)

				//清空m1 m2中相应的牌
				//1 先清空m1中的头一个牌(即是三带一中的单牌)
				if len(self.M1) >1{
					self.M1 = self.M1[1:]
				}

				//2 再清空 m1 中值等于m3_v的牌
				{
					idx := -1
					for i,v := range self.M1{
						if v == m3_v{
							idx = i
							break
						}
					}
					if idx != -1{
						self.M1 = append(self.M1[0:idx],self.M1[idx+1:]...)
					}
				}

				//3 再清空 m2中值等于m3_v的牌
				{
					idx := -1
					for i,v := range self.M2{
						if v == m3_v{
							idx = i
							break
						}
					}
					self.M2 = append(self.M2[0:idx],self.M2[idx+1:]...)
				}
			}else{
				//M1中没有值 就组成三牌
				node := &ThreeCarder{
					ThreeCard:m3_v,
				}
				self.ThreeCards = append(self.ThreeCards,*node)
			}
		}

		//M3已经拆完牌 清空
		self.M3 = self.M3[0:0]
	}
}

func (self* Parser) GetBeautifulCard(a []uint32){
	self.fillFourSourceList(a)
	self.SplitCard()
}

//填充4个辅助数组 作为拆牌的源数组
func (self* Parser) fillFourSourceList(a []uint32){

	self.Cards = self.Cards[0:0]
	self.M1 = self.M1[0:0]
	self.M2 = self.M2[0:0]
	self.M3 = self.M3[0:0]
	self.M4 = self.M4[0:0]

	//这个操作是为了利用int类型API做排序 begin
	for _,v := range a{
		self.Cards = append(self.Cards,int(v))
	}
	sort.Ints(self.Cards)
	fmt.Printf("排序后的手牌是:%+v\n",self.Cards)
	//end

	for _,v := range self.Cards{
		if self.check(self.M1,v) == false{
			//M1中不存在 插入
			self.M1 = append(self.M1,v)
			continue

		}else if(self.check(self.M2,v) == false){
			//M2中不存在 插入
			self.M2 = append(self.M2,v)
			continue

		}else if(self.check(self.M3,v) == false){
			//M3中不存在 插入
			self.M3 = append(self.M3,v)
			continue

		}else if(self.check(self.M4,v) == false){
			//M4中不存在 插入
			self.M4 = append(self.M4,v)
			continue
		}
	}
}

func (self* Parser) check(which []int,value int) bool {

	for _,v := range which{
		if v == value{
			return true
		}
	}
	return false
}

func (self* Parser) display(){
	fmt.Println()
	for _,v := range self.M1{
		fmt.Printf("%d,",v)
	}
	fmt.Println()
	for _,v := range self.M2{
		fmt.Printf("%d,",v)
	}
	fmt.Println()
	for _,v := range self.M3{
		fmt.Printf("%d,",v)
	}
	fmt.Println()
	for _,v := range self.M4{
		fmt.Printf("%d,",v)
	}
	fmt.Println()
}