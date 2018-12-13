package game

import (
	"trunk/cellnet"
)

var RegGameLogicFunction =  map[string]cellnet.EventCallback{}

func RegeditLogicProcessFunction(gamename string,f cellnet.EventCallback)  {
	RegGameLogicFunction[gamename] = f
}



