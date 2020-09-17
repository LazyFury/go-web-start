#! /bin/bash
rm -r ./dist/
# 创建文件夹
mkdir -p  dist/{config,wwwroot,template}
echo '编译项目中..'$1
export TempDir=./dist
export GOOS=''
export File=main
while [ -n "$1" ]  
do  
  case "$1" in   
    -linux)  
        echo "发现 -a 选项 编译目标：$1"  
        GOOS=linux
        File=linux-main
        ;;  
    -win)  
        echo "发现 -a 选项 编译目标：$1"  
        GOOS=windows
        File=main.exe
        ;;
    *)  
        echo "$1 is not an option,由go build选择当前默认编译平台"  
        ;;  
  esac  
  shift  
done

export CGO_ENABLED=0
export GOARCH=amd64

echo  GOOS=$GOOS   go build -o $TempDir/$File
eval  GOOS=$GOOS   go build -o $TempDir/$File


# 拷贝静态资源和配置文件
echo ' - 拷贝配置文件...'
cp -r config/config.json $TempDir/config/config.json
# cp -r config/database.db $TempDir/config/database.db
cp -r ../docker/zoneinfo.zip $TempDir/config/zoneinfo.zip

echo ' - 拷贝资源文件...'
cp -r wwwroot/* $TempDir/wwwroot
echo ' - 拷贝模板文件...'
cp -r template/* $TempDir/template
echo ' - 创建日志目录...'
mkdir $TempDir/log
echo '完成...'