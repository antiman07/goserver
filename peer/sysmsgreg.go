package peer

import (
	"trunk/cellnet"
	"trunk/cellnet/codec"
	_ "trunk/cellnet/codec/binary"
	"trunk/cellnet/util"
	"reflect"
)

func init() {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*cellnet.SessionAccepted)(nil)).Elem(),
		ID:    int(util.StringHash("cellnet.SessionAccepted")),
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*cellnet.SessionConnected)(nil)).Elem(),
		ID:    int(util.StringHash("cellnet.SessionConnected")),
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*cellnet.SessionConnectError)(nil)).Elem(),
		ID:    int(util.StringHash("cellnet.SessionConnectError")),
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*cellnet.SessionClosed)(nil)).Elem(),
		ID:    int(util.StringHash("cellnet.SessionClosed")),
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*cellnet.SessionCloseNotify)(nil)).Elem(),
		ID:    int(util.StringHash("cellnet.SessionCloseNotify")),
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*cellnet.SessionInit)(nil)).Elem(),
		ID:    int(util.StringHash("cellnet.SessionInit")),
	})
}
