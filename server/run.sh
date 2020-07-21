#!/bin/bash

pid=$(ps -ef |grep ./debug_bin |grep -v grep | awk '{print $2}');
echo "查找运行中 ./debug_bin 的进程..."

if [ ! -n "$pid" ]
then
  echo "没有运行中的进程";
else
  for i in `echo $pid`
    do
      echo "kill 进程pid $i 通知原进程fork子进程"
      kill $i
    done
fi
echo "building..."
go build -o debug_bin .
echo "restarting..."
nohup ./debug_bin >> ./log/nohup.log &
echo 'done.'