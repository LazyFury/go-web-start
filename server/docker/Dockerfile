FROM 'golang'

# Configure Go 生产模式不需要安装go
ENV GOPROXY http://goproxy.cn/
ENV GO111MODULE on
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/labstack/echo/v4
