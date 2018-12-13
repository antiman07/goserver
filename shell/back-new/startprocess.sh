#!/bin/bash
num=$1

start(){
	sh ./rmtxt.sh

	nohup ./rpclog &

	sleep 1s

	if [ $num -eq 0 ];then
		
		nohup ./myclient -b 1 -i 30000 &
	else 	
		nohup ./myclient -b $num -i 30000 &
	fi
}

start
