package ddz

import (
	"fmt"
	"sort"
	"trunk/cellnet/myrpc"
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

	//顺子牌 和 三牌 调用顺序有讲究,要充分考虑组装类型牌时 删除其他数组的牌
	self.fillThreeCard()
	self.fillDoubleLink() //组装多顺子牌时 把M2 M1中相应的牌删除掉

	self.fillSingleLink() //组装单顺子牌时 切记把M2中本组成对子的牌移到单张牌中


	self.fillDoubleCard()
	self.fillSingleCard()
}

//组装单牌 3 5 11 12 (通过拆火箭 炸弹 双顺 单顺 三牌 对牌后 M1中剩下的就都是单牌了)
func (self* Parser) fillSingleCard(){

	self.SingleCards = append(self.SingleCards,self.M1...)
	//排下顺序
	sort.Ints(self.SingleCards)
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
			self.M1 = append(self.M1[:del_index], self.M1[del_index+1:]...)
		}
	}

	//清空M2
	self.M2 = self.M2[0:0]

}

//组装火箭牌 从单牌M1 中检测是否有火箭(即是大小鬼 14 15)
func (self* Parser) fillRocket(){

	var it,jt int = -1,-1

	for i,v := range self.M1{

		if v == SMALL_GHOST{
			it = i
		}else if v == BIG_GHOST{
			jt = i
		}
	}

	if it != -1 && jt != -1{

		//大小鬼填充到火箭组
		self.Rocket = append(self.Rocket,self.M1[it],self.M1[jt])
		//remove M1 中的大小鬼
/*		self.M1 = append(self.M1[:it],self.M1[it+1:]...)
		self.M1 = append(self.M1[:jt],self.M1[jt+1:]...)*/

		host1 := self.M1[it]
		host2 := self.M1[jt]

		for i,v := range self.M1{
			if v == host1{
				self.M1 = append(self.M1[:i],self.M1[i+1:]...)
				break
			}
		}

		for i,v := range self.M1{
			if v == host2{
				self.M1 = append(self.M1[:i],self.M1[i+1:]...)
				break
			}
		}

	}
}

//组装炸弹牌 5555 9999
func (self* Parser) fillZhaDan(){
	for i := 0; i < len(self.M4); i++{

		zhadan := self.M4[i]

		self.Bombes = append(self.Bombes,zhadan)


		//remove M1 M2 M3 M4 中的炸弹牌
		//移除self.M1中的炸弹牌 begin
		var index = -1
		for i,value := range self.M1{
			if value == zhadan{
				index = i
				break
			}
		}

		if index != -1{
			self.M1 = append(self.M1[:index],self.M1[index+1:]...)
		}

		//移除self.M2中的炸弹牌 begin
		index = -1
		for i,value := range self.M2{
			if value == zhadan{
				index = i
				break
			}
		}

		if index != -1 {
			self.M2 = append(self.M2[:index], self.M2[index+1:]...)
		}


		//移除self.M3中的炸弹牌 begin
		index = -1
		for i,value := range self.M3{
			if value == zhadan{
				index = i
				break
			}
		}
		if index != -1 {
			self.M3 = append(self.M3[:index], self.M3[index+1:]...)
		}
		//end
	}

	//M4已经拆完牌 清空
	self.M4 = self.M4[0:0]
}

func (self* Parser) fillDoubleLink()  {
	//组装双顺牌(分析M2中是否有超过3张连续的牌 eg:112233 66778899)
	for{
		sub_series_arr := MaxSerialArr(self.M2)
		if nil == sub_series_arr{
			return
		}

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
						self.M1 = append(self.M1[:m1_i],self.M1[m1_i+1:]...)
						break
					}
				}

				for m2_i,m2_v := range self.M2{
					if m2_v == v{
						self.M2 = append(self.M2[:m2_i],self.M2[m2_i+1:]...)
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
		if nil == sub_series_arr{
			return
		}

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

				//1 移除M1中顺子牌
				for m1_i,m1_v := range self.M1{
					if m1_v == v{
						self.M1 = append(self.M1[:m1_i],self.M1[m1_i+1:]...)
						break
					}
				}

				//2 把本应该是对子的牌从M2中移除掉 放到M1中拆为单牌
				for m2_i,m2_v := range self.M2{
					if m2_v == v{
						self.M2 = append(self.M2[:m2_i],self.M2[m2_i+1:]...)
						self.M1 = append(self.M1,m2_v)
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

	for _,m3_v := range self.M3{

			//1 清空 m1 中值等于m3_v的牌
			{
				idx := -1
				for i,v := range self.M1{
					if v == m3_v{
						idx = i
						break
					}
				}
				if idx != -1{
					self.M1 = append(self.M1[:idx],self.M1[idx+1:]...)
				}
			}

			//2 清空 m2中值等于m3_v的牌
			{
				idx := -1
				for i,v := range self.M2{
					if v == m3_v{
						idx = i
						break
					}
				}
				if idx != -1{
					self.M2 = append(self.M2[:idx],self.M2[idx+1:]...)
				}
			}

			//3 在M1中找到单牌组成三带一组合(即是遍历M1数值,不存在于M3和M2中,就认为是单牌了)
			{
				idx := -1

				for m1_i,m1_v := range self.M1{

					find := false

					for _,m3_v := range self.M3{
						if m1_v == m3_v {
							find = true
							break
						}
					}

					//在M3中找到了就没必要再M2中找了,直接continue
					if find == true{
						continue
					}

					for _,m2_v := range self.M2{
						if m1_v == m2_v{
							find = true
							break
						}
					}


					if find == false{
						//M1中有 且M3中没有 此时认成单张牌
						idx = m1_i
						break
					}
				}


				//生出三带一并移除单张牌
				if idx != -1{
					node := &ThreeOneCarder{
						ThreeCard:m3_v,
						SingleCard:self.M1[idx],//这个确实是单牌(仅仅存在M1中,M2 M3中都没这个值)
					}
					self.ThreeOneCards = append(self.ThreeOneCards,*node)

					self.M1 = append(self.M1[:idx],self.M1[idx+1:]...)
				}else{
					//只有三张一样的牌 没有单张 (这种情况很少见)
					node := &ThreeCarder{
						ThreeCard:m3_v,
					}
					self.ThreeCards = append(self.ThreeCards,*node)
				}
			}
	}

	//M3已经拆完牌 清空
	self.M3 = self.M3[0:0]
}

func (self* Parser) GetBeautifulCard(a []uint32){
	self.fillFourSourceList(a)
	self.SplitCard()
	self.display()
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
	myrpc.Rpcqueue <- fmt.Sprintf("\n手牌排序: %+v",self.Cards)
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
	myrpc.Rpcqueue <- fmt.Sprintf("\n拆分成四个辅助数组: %+v,%+v,%+v,%+v",self.M1,self.M2,self.M3,self.M4)
}

func (self* Parser) check(which []int,value int) bool {

	for _,v := range which{
		if v == value{
			return true
		}
	}
	return false
}

//打印拆好的牌
func (self* Parser) display(){

	myrpc.Rpcqueue <- fmt.Sprintf("\n+++++++++拆好的牌数据 begin+++++++++++++++++++:\n")
	myrpc.Rpcqueue <- "火箭牌:"
	if len(self.Rocket) != 0{
		myrpc.Rpcqueue <- "大鬼 小鬼"
	}

	myrpc.Rpcqueue <- fmt.Sprintf("\n炸弹牌:")
	for _,v := range self.Bombes{
		myrpc.Rpcqueue <- fmt.Sprintf("%d%d%d%d\n",v,v,v,v)
	}
	myrpc.Rpcqueue <- fmt.Sprintf("\n双顺牌:")
	for _,v := range self.DoubleLink{

		begin_card := v.BeginValue
		var ll []int
		var i int

		for i = 0; i < v.Len; i++{
			ll = append(ll, begin_card + i)
			ll = append(ll, begin_card + i)
		}
		myrpc.Rpcqueue <- fmt.Sprintf("%+v\n",ll)

	}
	myrpc.Rpcqueue <- fmt.Sprintf("\n单顺牌:")
	for _,v := range self.SingleLink{

		begin_card := v.BeginValue
		var ll []int
		var i int

		for i = 0; i < v.Len; i++{
			ll = append(ll, begin_card + i)
		}
		myrpc.Rpcqueue <- fmt.Sprintf("%+v\n",ll)

	}

	myrpc.Rpcqueue <- fmt.Sprintf("\n三带一牌:")
	for _,v := range self.ThreeOneCards{
		myrpc.Rpcqueue <- fmt.Sprintf("%d%d%d-%d\n",v.ThreeCard,v.ThreeCard,v.ThreeCard,v.SingleCard)
	}

	myrpc.Rpcqueue <- fmt.Sprintf("\n三牌:")
	for _,v := range self.ThreeCards{
		myrpc.Rpcqueue <- fmt.Sprintf("%d%d%d\n",v.ThreeCard,v.ThreeCard,v.ThreeCard)
	}

	myrpc.Rpcqueue <- fmt.Sprintf("\n对子牌:")
	for _,v := range self.DoubleCards{
		myrpc.Rpcqueue <- fmt.Sprintf("%d%d  ",v,v)
	}

	myrpc.Rpcqueue <- fmt.Sprintf("\n单张牌:")
	for _,v := range self.SingleCards{
		myrpc.Rpcqueue <- fmt.Sprintf("%d  ",v)
	}

	myrpc.Rpcqueue <- fmt.Sprintf("\n+++++++++拆好的牌数据 end+++++++++++++++++++:")
}