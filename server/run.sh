#!/bin/bash
tmp_path=""

export CGO_ENABLED=0
export GOARCH=amd64
# export GOOS=

if [ ! -n "$1" ]
then
  echo "请输入程序路径 like ./debug_bin"
  exit
fi


tmp_path=$1
echo "building..."
go build -o $tmp_path

pid=$(ps -ef |grep $tmp_path |grep -v grep | grep -v 'run.sh' |  awk '{print $2}');
echo "查找运行中 $tmp_path 的进程..."

if [ ! -n "$pid" ]
then
  echo "没有运行中的进程";
  echo "restarting..."
  # nohup  $tmp_path >> ./log/nohup.log &
  exec $tmp_path >> ./log/nohup.log &
  #通知重启一下  结束当前回话 nohup有些异常
  exec $0 $1
else
  for i in `echo $pid`
    do
      echo "kill 进程pid $i 通知原进程fork子进程"
      kill -s 1 $i
    done
fi


echo 'done.'