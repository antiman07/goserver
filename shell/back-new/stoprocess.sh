ps aux|grep rpclog|grep -v grep|awk '{print $2}'|xargs -i kill {}
ps aux|grep myclient|grep -v  grep|awk '{print $2}'|xargs -i kill {}
#不能用kill -9,不然GO进程接收不到信号
