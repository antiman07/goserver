package proc

import (
	"trunk/cellnet"
	"sort"
)

type ProcessorBinder func(bundle ProcessorBundle, userCallback cellnet.EventCallback)

var (
	procByName = map[string]ProcessorBinder{}
)

// 注册事件处理器，内部及自定义收发流程时使用
func RegisterProcessor(procName string, f ProcessorBinder) {

	procByName[procName] = f
}

// 获取处理器列表
func ProcessorList() (ret []string) {

	for name := range procByName {
		ret = append(ret, name)
	}

	sort.Strings(ret)
	return
}

// 绑定固定回调处理器, procName来源于RegisterProcessor注册的处理器，形如: 'tcp.ltv'
//参数类型 cellnet.EventCallback  见 cellnet\processor.go 定义
func BindProcessorHandler(peer cellnet.Peer, procName string, userCallback cellnet.EventCallback) {
    //下面代码 参照 tcp\setup.go init函数 便于理解
	if proc, ok := procByName[procName]; ok {

		bundle := peer.(ProcessorBundle) //入参 peer== tcpAcceptor struct === tcpSession struct {包含 ProcessorBundle}

		proc(bundle, userCallback)//通过函数proc 把userCallback绑定到 实例bundle上 见

	} else {
		panic("processor not found:" + procName)
	}
}
