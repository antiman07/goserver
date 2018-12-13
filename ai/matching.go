package ai

import (
	"time"
	"trunk/cellnet/util"
)


var Matching_queue1 util.Queue

/*var Matching_queue2 []*SessionInfo
var Matching_queue3 []*SessionInfo

var Matched_queue1 []*SessionInfo
var Matched_queue2 []*SessionInfo
var Matched_queue3 []*SessionInfo*/

func Matching_1_loop() {

	for {
		if Matching_queue1.Count() % 3 == 0 {
			//人数够 就开桌
			for i := 0; i < 3; i++ {
				s1 := Matching_queue1.Dequeue()
				s2 := Matching_queue1.Dequeue()
				s3 := Matching_queue1.Dequeue()
				go EnterRoom(1,s1,s2,s3)
			}

		} else {
			time.Sleep(time.Second * 5)
		}
	}

}

/*func Matched_1_loop(){
	for{
		//准备发牌数据,发给玩家
	}
}*/