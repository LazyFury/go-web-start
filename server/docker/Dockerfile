FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update
# RUN apt install -y ffmpeg
# RUN apt install lame
# Configure Go 生产模式不需要安装go
ENV GOROOT /usr/local/go
ENV GOPATH /root/go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH
ENV GOPROXY http://goproxy.cn/
ENV GO111MODULE on
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin


# 二进制文件直接运行 not fund 
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
# 时区文件
COPY  ./zoneinfo.zip /opt/zoneinfo.zip
ENV ZONEINFO /opt/zoneinfo.zip
EXPOSE 8080