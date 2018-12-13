package proc

import (
	"trunk/cellnet"
)

// 处理器设置接口，由各Peer实现
type ProcessorBundle interface {

	// 设置 传输器，负责收发消息
	SetTransmitter(v cellnet.MessageTransmitter)

	// 设置 接收后，发送前的事件处理流程
	SetHooker(v cellnet.EventHooker)

	// 设置 接收后最终处理回调
	SetCallback(v cellnet.EventCallback)
}

// 放队列中回调
//这种写法有点绕 首先函数 入参和返回类型一样,都是type EventCallback func(ev Event),返回的是函数地址
//其次 这个函数是闭包函数,子函数 if 里的 callback 引用的是外层函数的 callback
func NewQueuedEventCallback(callback cellnet.EventCallback) cellnet.EventCallback {
//procbundle.go PostEvent 函数中 self.callback(ev) 调用的就是下面的代码
	return func(ev cellnet.Event) {
		if callback != nil {
			cellnet.SessionQueuedCall(ev.Session(), func() {

				callback(ev)
			})
		}
	}
}

// 当需要多个Hooker时，使用NewMultiHooker将多个hooker合并成1个hooker处理
type MultiHooker []cellnet.EventHooker

func (self MultiHooker) OnInboundEvent(input cellnet.Event) (output cellnet.Event) {

	for _, h := range self {

		input = h.OnInboundEvent(input)
	}

	return input
}

func (self MultiHooker) OnOutboundEvent(input cellnet.Event) (output cellnet.Event) {

	for _, h := range self {

		input = h.OnOutboundEvent(input)
	}

	return input
}

func NewMultiHooker(h ...cellnet.EventHooker) cellnet.EventHooker {

	return MultiHooker(h)
}
