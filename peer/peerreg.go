package peer

import (
	"trunk/cellnet"
	"sort"
)

type PeerCreateFunc func() cellnet.Peer

var creatorByTypeName = map[string]PeerCreateFunc{}

// 注册Peer创建器 tcp udp http...
func RegisterPeerCreator(f PeerCreateFunc) {

	// 临时实例化一个，获取类型
	dummyPeer := f()

	if _, ok := creatorByTypeName[dummyPeer.TypeName()]; ok {
		panic("Duplicate peer type")
	}

	creatorByTypeName[dummyPeer.TypeName()] = f
}

// Peer创建器列表
func PeerCreatorList() (ret []string) {

	for name := range creatorByTypeName {
		ret = append(ret, name)
	}

	sort.Strings(ret)
	return
}

// 创建一个Peer
func NewPeer(peerType string) cellnet.Peer {
	peerCreator := creatorByTypeName[peerType]
	if peerCreator == nil {
		panic("Peer type not found: " + peerType)
	}

	return peerCreator()
}

// 创建Peer后，设置基本属性
func NewGenericPeer(peerType, name, addr string, q cellnet.EventQueue) cellnet.GenericPeer {

	p := NewPeer(peerType) //根据入参 peerType 得知创建的是 tcp\acceptor.go tcpAcceptor实例对象
	// p是tcpAcceptor结构的实例,该结构中的peer.CorePeerProperty实现了GenericPeer接口的所有方法,故可以这样转换.
	gp := p.(cellnet.GenericPeer)
	gp.SetName(name)
	gp.SetAddress(addr)
	gp.SetQueue(q)
	return gp
}
