package ai

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var Paidata = [5][13]int {
	{0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x09,0x0a,0x0b,0x0c,0x0d},//方块(Diamond)
	{0x11,0x12,0x13,0x14,0x15,0x16,0x17,0x18,0x19,0x1a,0x1b,0x1c,0x1d},//梅花(Club)
	{0x21,0x22,0x23,0x24,0x25,0x26,0x27,0x28,0x29,0x2a,0x2b,0x2c,0x2d},//红桃(Heart)
	{0x31,0x32,0x33,0x34,0x35,0x36,0x37,0x38,0x39,0x3a,0x3b,0x3c,0x3d},//黑桃(Spade)
	{0x4E,0x4F}, //王(Joker)
}

var paidata = []int{
	1,2,3,4,5,6,7,8,9,10,11,12,13,
	1,2,3,4,5,6,7,8,9,10,11,12,13,
	1,2,3,4,5,6,7,8,9,10,11,12,13,
	1,2,3,4,5,6,7,8,9,10,11,12,13,
	14,15,
}

func fapai() interface{} {
	flag := [54]int{}
	rand.Seed(time.Now().Unix())

	result := make(map[int][]int)
	for i := 0; i < 3; i++ {
		rst := make([]int, 0)

		for index := 0; index < 17; index++ {
			m := rand.Intn(54)
			if flag[ m ] == 1{
				index--
				continue
			}

			flag[m] = 1
			rst = append(rst, paidata[m])
		}
		sort.Ints(rst)
		fmt.Printf("%d",rst )
		fmt.Println()
		result[i] = rst
	}
	for k,_ := range flag{
		if flag[k] == 0{
			fmt.Printf("底牌是 %d\n",paidata[k])
		}
	}
	fmt.Println(flag)
	return result
}

func EnterRoom(whichroom int,s1 interface{},s2 interface{},s3 interface{}) error{
	if whichroom >3 || whichroom <=0{
		return errors.New("房间不存在,只有1 2 3 房间")
	}
/*	session0 := s1.(cellnet.Session)
	session1 := s2.(cellnet.Session)
	session2 := s3.(cellnet.Session)
	switch whichroom {
	case 1:
		rst := get_free_desk_from_room1()
		if rst != nil{
			desk := rst.(*Desk)
			desk.IsGameing = true
			desk.ID = 1000
			desk.CurPlay = ""
			pailist := fapai().(map[int][]int)
			pson0 := desk.PeopleList[0]
			pson0.MySelfName = session0.GetPlayer().(role.Player).Att.Account
			pson0.MyHandPai = pailist[0]
			pson0.Ses = session0
			pson0.Sesslist = append(pson0.Sesslist,session1)
			pson0.Sesslist = append(pson0.Sesslist,session2)
			session0.GetPlayer().(role.Player).Desk = desk
			//session0.Send()

			pson1 := desk.PeopleList[1]
			pson1.MySelfName = session1.GetPlayer().(role.Player).Att.Account
			pson1.MyHandPai = pailist[1]
			pson1.Ses = session1
			pson1.Sesslist = append(pson0.Sesslist,session0)
			pson1.Sesslist = append(pson0.Sesslist,session2)
			session1.GetPlayer().(role.Player).Desk = desk

			pson2 := desk.PeopleList[2]
			pson2.MySelfName = session2.GetPlayer().(role.Player).Att.Account
			pson2.MyHandPai = pailist[2]
			pson2.Ses = session2
			pson2.Sesslist = append(pson0.Sesslist,session0)
			pson2.Sesslist = append(pson0.Sesslist,session1)
			session2.GetPlayer().(role.Player).Desk = desk

			fmt.Println(desk.ID,desk.IsGameing,desk.PeopleList)
			for _,node := range desk.PeopleList{
				fmt.Println(node.MySelfName,node.MyHandPai)
			}
		}else{
			return errors.New("房间没空余牌桌")
		}

	default:
		return errors.New("选择房间出错")
	}*/

	return nil
}


