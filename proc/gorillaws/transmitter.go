package gorillaws

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"trunk/cellnet"
	"trunk/cellnet/codec"
	"trunk/cellnet/game"
	"trunk/cellnet/myrpc"
)

const (
	MsgIDSize = 2 // uint16
)


type WSMessageTransmitter struct {
}

func (WSMessageTransmitter) OnRecvMessage(ses cellnet.Session) (msg interface{}, err error) {

	conn, ok := ses.Raw().(*websocket.Conn)

	// 转换错误，或者连接已经关闭时退出
	if !ok || conn == nil {
		return nil, errors.New("OnRecvMessage conn is nil")
	}

	var messageType int
	var raw []byte
	messageType, raw, err = conn.ReadMessage()

	if err != nil{
		myrpc.Rpcqueue <- fmt.Sprintf("recv interface error1=%v",err.Error())
		return nil,err
	}

	switch messageType {
	case websocket.BinaryMessage:
		metadataID := binary.BigEndian.Uint16(raw)
		cmdid := 0
		msgData := raw[MsgIDSize:]

		//适应多款游戏 ID转换 多GO程读取 这里不加锁不知道会不会出错 begin
		for k,v := range game.ConvertMsgID {
			if v == metadataID{
				cmdid = k
			}
		}
		//end

		msg,_,err = codec.DecodeMessage(cmdid,msgData)
	}
	return
}


func (WSMessageTransmitter) OnSendMessage(ses cellnet.Session, msg interface{}) error {

	conn, ok := ses.Raw().(*websocket.Conn)

	// 转换错误，或者连接已经关闭时退出
	if !ok || conn == nil {
		return errors.New("error conn is nil")
	}

	var(
		msgData []byte
		//msgID int
		cmd uint16
	)
	switch m := msg.(type){
	case *cellnet.RawPacket:
		msgData = m.MsgData
		//msgID = m.MsgID
	default:
		var err error
		var meta *cellnet.MessageMeta
		msgData,meta,err = codec.EncodeMessage(msg,nil)
		if err != nil{
			return err
		}
		//适应多款游戏 ID转换 begin 多GO程读取 begin
		cmd = game.ConvertMsgID[meta.ID] //仅仅读取 不用加锁
		//end
	}

	pkt := make([]byte,MsgIDSize+len(msgData))
	binary.BigEndian.PutUint16(pkt,cmd)
	copy(pkt[MsgIDSize:],msgData)
	err := conn.WriteMessage(websocket.BinaryMessage,pkt)
	if err != nil{
		myrpc.Rpcqueue <- fmt.Sprintf("send error =%v",err.Error())
		return err
	}
	return nil
}
