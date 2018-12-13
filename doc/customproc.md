# 定制Processor
cellnet提供的Processor能满足基本的通信封包格式的处理，但在特殊封包定制需求时，仍然需要编写自己的Processor

**定制Processor的根本目的：复用业务逻辑层与协议通信层之间的逻辑**

例如：
- 编写服务器逻辑时，已有Unity3D客户端网络库及封包格式，需要服务器与客户端通信，封包格式为私有协议。

- 在连接建立后，需要有握手过程（加密秘钥交换，对时等），此时将这个过程封装到Processor中，可以方便Peer的快速复用。

- 服务器互相连接时，需要标识连接方的服务类型。

- KCP协议需要建立在UDP协议层，UDP收发过程已由Peer部分完成，KCP协议解析只需要放在Processor即可。

- RPC（远程过程调用）是建立在tcp协议之上的调用封装，也可以使用Processor来完成封装过程。

- 消息收发量统计。

- 使用nsq、mysql的Peer的基础上，还需要对用户消息进行二次封装，以实现扩展消息。

## 不需要使用Processor扩展的情况

- 编码层更换

   从二进制编码更换为ProtocolBuffer编码，从JSON更换为二进制编码等。这种情况下直接使用codec编码包即可。

- 具体的业务逻辑


## 内建处理器(tcp.ltv)封包格式

tcp.ltv使用util/packet.go中的函数解析封包，同时处理粘包问题，封包格式如下：

功能 | 类型 | 备注
---|---|---
包体大小 | uint16 | 不包含包大小自己，例如，本字段为10时，包体还需要读10字节
包体中的消息ID | uint16 | 包.消息名 的hash值(util.StringHash)
包体中的用户消息数据 | []byte | 用户的消息大小，需要使用codec包解码

封包解析请参考:
https://github.com/cellnet/blob/master/proc/tcp/transmitter.go


## 内建处理器(udp.ltv)封包格式

注意UDP封包总长度不超过MTU

功能 | 类型 | 备注
---|---|---
包体大小 | uint16 | 只做UDP包完整性验证,包含包体大小本身
包体中的消息ID | uint16 | 包.消息名 的hash值(util.StringHash)
包体中的用户消息数据 | []byte | 用户的消息大小，需要使用codec包解码

封包解析请参考:
https://github.com/cellnet/blob/master/proc/udp/recv.go
