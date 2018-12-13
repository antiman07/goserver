package ai

type People struct{
	IsDizhu bool
	MySelfName string
	MyHandPai []int
	//Ses cellnet.Session
	//Sesslist []cellnet.Session
}

type Desk struct{
	ID int   //桌号
	IsGameing bool  //是否正在游戏中
	CurPlay string  //当前轮到谁出牌
	PeopleList []*People

}

