FROM alpine:latest

RUN apk add --no-cache go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin


# Install Glide

WORKDIR  /run-server
# RUN git clone https://github.com/Treblex/go-echo-demo
WORKDIR  /run-server/go-echo-demo
COPY . .
RUN go build -o main
EXPOSE 8080
ENTRYPOINT ["./main"]