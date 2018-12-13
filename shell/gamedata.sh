#!/bin/bash
whichplayer=$1

comm(){
	OnlineUserNums=`grep -rn "connect success" ./log.txt|wc -l`
	echo "并发在线玩家数量:$OnlineUserNums"
	echo "目前实时在线玩家数量`netstat -an|grep 11303|wc -l`"
	allnums=`grep -rn "当前金币数量" ./log.txt|wc -l`
	echo "所有玩家合计玩了$allnums局游戏"
}

player(){
	if [ -z "$whichplayer" ];then
		while((i<=10))
		do
			let tt=20000+i
			let i+=1
			tmp="玩家acc_$tt"
		    	nums=`grep -rn $tmp ./log.txt |grep "当前金币数量"|wc -l`
                	echo "$tmp 目前玩了$nums局游戏"
		done
	else
		nums=`grep -rn $whichplayer ./log.txt |grep "当前金币数量"|wc -l`
        	echo "$whichplayer 目前玩了$nums局游戏">>./stat-result.txt
	fi	
}
lasttime(){
	if [ -z "$whichplayer" ];then
		while((i<=100))
		do
			let tt=10000+i
			let i+=1
			tmp="玩家acc_$tt"
			grep -rn $tmp ./log.txt |grep "当前金币数量"|tail -1>>./lasttime
		done
	fi
}

comm
player
#lasttime
