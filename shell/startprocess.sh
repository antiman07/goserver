#!/bin/bash
start(){
	if [ ! "$1" ];then
	 robots=3000
	else robots=$1
	fi

	nohup ./rpclog &
	sleep 1s
	nohup ./myclient -b $robots -i 20000 &
}

start
