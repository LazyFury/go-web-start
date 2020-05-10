FROM gitpod/workspace-mysql
                    
USER suke

# Install custom tools, runtime, etc. using apt-get
# For example, the command below would install "bastet" - a command line tetris clone:
#
# RUN sudo apt-get -q update && #     sudo apt-get install -yq bastet && #     sudo rm -rf /var/lib/apt/lists/*
#
# More information: https://www.gitpod.io/docs/config-docker/
RUN apk update
RUN apk add --no-cache go
# Configure Go
ENV GOPATH /workspace/go
ENV PATH $GOPATH/bin:$PATH
# ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

COPY  ./server/config/zoneinfo.zip /opt/zoneinfo.zip
ENV ZONEINFO /opt/zoneinfo.zip
WORKDIR /workspace/go-echo-demo/server
EXPOSE 8080