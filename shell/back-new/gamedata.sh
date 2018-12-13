#!/bin/bash
whichplayer=$1
commdata(){
	OnlineUserNums=`grep -rn "connect success" ./log.txt|wc -l`
	echo "并发在线玩家数量:$OnlineUserNums"
	echo "目前实时在线玩家数量`netstat -an|grep 11303|wc -l`"
	allnums=`grep -rn "当前金币数量" ./log.txt|wc -l`
	echo "所有玩家合计玩了$allnums局游戏"
}

calc_all_player(){
	
	cat ./log.txt |awk -F" " '$6~/acc/ {print $6}'|sort -k1 -n|uniq -c|awk '{print "玩家"$2,"玩了"$1"局"}'>allplayerdata.txt
}

player(){
       	echo "$whichplayer 目前玩了 `grep -rn $whichplayer ./log.txt |grep "当前金币数量"|wc -l` 局游戏"
}

lasttime(){
	cat ./log.txt|awk '{print $3,$4,$5,$6,$7,$8}'|sort -k6 -n|uniq -c
}
	

commdata
calc_all_player
#lasttime
