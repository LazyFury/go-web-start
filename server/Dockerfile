FROM alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update
RUN apk add --no-cache go
# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
RUN apk add ffmpeg
RUN apk add lame

COPY  ./config/zoneinfo.zip /opt/zoneinfo.zip
COPY ./ /root/app1/
WORKDIR /root/app1
RUN go get -u
WORKDIR /root/app
ENV ZONEINFO /opt/zoneinfo.zip
EXPOSE 8080
CMD ["./linux-main"]