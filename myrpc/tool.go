package myrpc

import (
	"fmt"
	"net/rpc"
)


var rpcclient *rpc.Client
var Rpcqueue chan interface{}

type Arg struct{
	Logdata string
}

func Rpclog(log_addr string,whichtype int){

	client, err := rpc.DialHTTP("tcp",log_addr)
	if err != nil{
		fmt.Println(err.Error())
	}
	rpcclient = client

	ch := make(chan interface{},1000)

	Rpcqueue = ch

	reply := ""
	for node := range Rpcqueue{
		if node == nil{
			break
		}else{
			if whichtype == 1{
				rpcclient.Call("Logtxt.ServerWriteLog",Arg{Logdata:node.(string)},&reply)
			}else{
				rpcclient.Call("Logtxt.ClientWriteLog",Arg{Logdata:node.(string)},&reply)
			}

		}
	}
}