FROM alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update
RUN apk add fish
RUN apk add busybox-extras
RUN apk add --no-cache go
RUN apk add ffmpeg
RUN apk add lame

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /root/go
ENV PATH $GOPATH/bin:$PATH
ENV GOPROXY https://goproxy.cn
ENV GO111MODULE off
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

# 时区文件
COPY  ./config/zoneinfo.zip /opt/zoneinfo.zip
ENV ZONEINFO /opt/zoneinfo.zip
EXPOSE 8080