package game

// protobuf 脚本自动生成的协议字段ID 跟 具体游戏服务器指定的命令ID做映射 收发协议的头部要转换这2个字段
var ConvertMsgID map[int]uint16 //int:cellnet脚本自动生成的消息ID uint16:游戏服务器指定的消息ID

func init(){
	ConvertMsgID = make(map[int]uint16)
}
