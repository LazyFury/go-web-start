#!/bin/zsh

pid=$(ps -ef |grep ./debug_bin |grep -v grep | awk '{print $2}');
echo "pid = [$pid]"

if [ ! -n "$pid" ]
then
  echo "没有运行中的进程";
  go build -o debug_bin .;./debug_bin
fi

kill 1 "$pid";
go build -o debug_bin .;./debug_bin