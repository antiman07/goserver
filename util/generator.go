package util

import "sync/atomic"

var cur_largest *int64


func SetFrontRobotID(id int64){
	cur_largest = &id
}

//原子安全  递增生成整数 用于生成机器人登陆的账号
func GenIntValue() int64 {
	return atomic.AddInt64(cur_largest,1)
}
