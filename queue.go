package cellnet

import (
	"runtime/debug"
	"sync"
	"trunk/cellnet/myrpc"
)

// 事件队列
type EventQueue interface {

	// 事件队列开始工作
	StartLoop() EventQueue

	// 停止事件队列
	StopLoop() EventQueue

	// 等待退出
	Wait()

	// 投递事件, 通过队列到达消费者端
	Post(callback func())

	// 是否捕获异常
	EnableCapturePanic(v bool)
}

type eventQueue struct {
	queue chan func()

	endSignal sync.WaitGroup

	capturePanic bool
}

// 启动崩溃捕获
func (self *eventQueue) EnableCapturePanic(v bool) {
	self.capturePanic = v
}

// 派发事件处理回调到队列中
func (self *eventQueue) Post(callback func()) {

	if callback == nil {
		return
	}

	self.queue <- callback
}

// 保护调用用户函数
func (self *eventQueue) protectedCall(callback func()) {

	if self.capturePanic {
		defer func() {

			if err := recover(); err != nil {

				debug.PrintStack()
			}

		}()
	}

	callback()
}

// 开启事件循环
func (self *eventQueue) StartLoop() EventQueue {

	self.endSignal.Add(1)

	go func() {

		for callback := range self.queue {

			if callback == nil {
				myrpc.Rpcqueue <- "退出事件循环GO程"
				break
			}

			self.protectedCall(callback)
		}

		self.endSignal.Done()
	}()

	return self
}

// 停止事件循环
func (self *eventQueue) StopLoop() EventQueue {
	self.queue <- nil
	return self
}

// 等待退出消息
func (self *eventQueue) Wait() {
	self.endSignal.Wait()
}

const DefaultQueueSize = 100

// 创建默认长度的队列
func NewEventQueue() EventQueue {

	return &eventQueue{
		queue: make(chan func(), DefaultQueueSize),
	}
}

// 在会话对应的Peer上的事件队列中执行callback，如果没有队列，则马上执行
func SessionQueuedCall(ses Session, callback func()) {
	if ses == nil {
		return
	}
	q := ses.Peer().(interface {
		Queue() EventQueue
	}).Queue()

	QueuedCall(q, callback)
}

// 有队列时队列调用，无队列时直接调用
func QueuedCall(queue EventQueue, callback func()) {
	if queue == nil {
		callback()
	} else {
		queue.Post(callback)
	}
}
