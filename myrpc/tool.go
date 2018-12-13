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

func Rpclog(log_addr string){

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
			rpcclient.Call("Logtxt.WriteLog",Arg{Logdata:node.(string)},&reply)
		}
	}
}