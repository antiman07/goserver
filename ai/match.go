package ai

import (
)

func Build() *Matching{

	return &Matching{
		queue:NewDDZPipe(),
	}
}

//匹配队列
type Matching struct{
	queue *DDZPipe
}

//放到匹配队列中
func (self* Matching) Add(i interface{})  {
	self.queue.Add(i)
}

//退出匹配GO程
func (self* Matching) StopLoop() {
	self.Add(nil)
}

//开启匹配循环
func (self* Matching) StartLoop()  {

	var readyList []interface{}

	for {
		readyList = readyList[0:0]
		exit := self.queue.Pick(&readyList)
		if exit{
			break
		}else {
			go EnterRoom(readyList[0],readyList[1],readyList[2])
		}
	}
}



