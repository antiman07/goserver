#!/bin/bash

echo "连接数:`grep -rn "建链成功" ./clog.txt |wc -l`"
echo "释放连接数:`grep -rn "释放链接" ./clog.txt |wc -l`"

echo "退出玩家发送数据GO程 `grep -rn "退出发送数据GO程" ./clog.txt |wc -l`"

echo "退出玩家接收数据GO程 `grep -rn "退出接收数据GO程" ./clog.txt |wc -l`"

echo "退出心跳GO程 `grep -rn "退出心跳GO程" ./clog.txt |wc -l`"

echo "退出事件循环GO程数量:`grep -rn "退出事件循环GO程" ./clog.txt |wc -l`"
