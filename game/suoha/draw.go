package suoha

import (
	"fmt"
	"strings"
	"trunk/cellnet"
	"trunk/cellnet/game"
)

func Init_game_protocol(){

	// protobuf 脚本自动生成的协议字段ID 跟 具体游戏服务器指定的命令ID做映射 收发协议的头部要转换这2个字段
	metaByFullName := cellnet.GetMetaMap()

	protocolmap := ProtocolMap

	for name,meta := range metaByFullName {
		metadata := meta //这句话必须要 赋值给临时变量 网上搜range陷阱
		protocolname := name[strings.Index(name,".")+1 : len(name)]
		if msgid,ok := protocolmap[protocolname];ok{
			fmt.Printf("msgid=%d,metaid=%d,msg=%s\n",msgid,metadata.ID,protocolname)
			game.ConvertMsgID[metadata.ID] =  msgid //映射 自动生成的meteID ---> 服务器端指定的ID
		}else{
			fmt.Printf("%s protocol not regedit\n",protocolname)
		}
	}
}