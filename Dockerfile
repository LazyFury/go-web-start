FROM golang:latest

WORKDIR /usr/src/app
COPY . .

#go构建可执行文件
RUN go build .
#暴露端口
EXPOSE 80
#最终运行docker的命令
ENTRYPOINT  ["./dist/main-linux"]

