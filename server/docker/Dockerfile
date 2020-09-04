FROM debian:9.0

RUN printf "deb http://mirrors.aliyun.com/debian/ stretch main non-free contrib\ndeb-src http://mirrors.aliyun.com/debian/ stretch main non-free contrib\ndeb http://mirrors.aliyun.com/debian-security stretch/updates main\ndeb-src http://mirrors.aliyun.com/debian-security stretch/updates main\ndeb http://mirrors.aliyun.com/debian/ stretch-updates main non-free contrib\ndeb-src http://mirrors.aliyun.com/debian/ stretch-updates main non-free contrib\ndeb http://mirrors.aliyun.com/debian/ stretch-backports main non-free contrib\ndeb-src http://mirrors.aliyun.com/debian/ stretch-backports main non-free contrib\n" > /etc/apt/sources.list

# alpine
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# RUN apk update
# RUN apk add bash
# RUN apt install -y ffmpeg
# RUN apt install lame




RUN apt update &&  apt -y full-upgrade
RUN apt-get install -y procps && apt-get install -y wget 


RUN wget https://dl.google.com/go/go1.15.1.linux-amd64.tar.gz
RUN tar -xvf  go1.15.1.linux-amd64.tar.gz

# goproxy ca证书
RUN wget http://cacerts.digitalcertvalidation.com/TrustAsiaTLSRSACA.crt
RUN cp  TrustAsiaTLSRSACA.crt /usr/share/ca-certificates/goproxy.crt
RUN chmod a+r /usr/share/ca-certificates/goproxy.crt
RUN update-ca-certificates

# Configure Go 生产模式不需要安装go
ENV GOROOT /go
ENV GOPATH /root/go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH
ENV GOPROXY http://goproxy.cn/
ENV GO111MODULE on
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin


# 二进制文件直接运行 not fund 
# RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
# 时区文件
COPY  ./zoneinfo.zip /opt/zoneinfo.zip
ENV ZONEINFO /opt/zoneinfo.zip
EXPOSE 8080