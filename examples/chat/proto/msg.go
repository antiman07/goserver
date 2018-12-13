package proto

import (
	"fmt"
	"github.com/cellnet"
	"github.com/cellnet/codec"

	// 使用binary协议，因此匿名引用这个包，底层会自动注册
	_ "github.com/cellnet/codec/binary"
	_ "github.com/cellnet/codec/json"

	"github.com/cellnet/util"
	"reflect"
)

type ChatREQ struct {
	Content string
}

type ChatACK struct {
	Content string
	Id      int64
}

type TestEchoACK struct {
	Msg   string
	Value int32
}

type User_Info struct {
	UserId int32
	UserName string
	Chip int32
}

func (self *User_Info) String() string{ return fmt.Sprintf("%+v", *self)}
func (self *TestEchoACK) String() string { return fmt.Sprintf("%+v", *self) }
// 用于消息日志打印消息内容
func (self *ChatREQ) String() string { return fmt.Sprintf("%+v", *self) }
func (self *ChatACK) String() string { return fmt.Sprintf("%+v", *self) }

// 引用消息时，自动注册消息，这个文件可以由代码生成自动生成
func init() {

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*ChatREQ)(nil)).Elem(),
		ID:    int(util.StringHash("proto.ChatREQ")),
	})

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*ChatACK)(nil)).Elem(),
		ID:    int(util.StringHash("proto.ChatACK")),
	})

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("json"),
		Type:  reflect.TypeOf((*TestEchoACK)(nil)).Elem(),
		ID:    int(util.StringHash("main.TestEchoACK")),
	})

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec:codec.MustGetCodec("json"),
		Type:reflect.TypeOf((*User_Info)(nil)).Elem(),
		ID:int(util.StringHash("main.User_Info")),
	})

}
