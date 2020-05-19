#-------------------------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See https://go.microsoft.com/fwlink/?linkid=2090316 for license information.
#-------------------------------------------------------------------------------------------------------------

# To fully customize the contents of this image, use the following Dockerfile instead:
# https://github.com/microsoft/vscode-dev-containers/tree/v0.117.1/containers/ubuntu-18.04-git/.devcontainer/Dockerfile
FROM mcr.microsoft.com/vscode/devcontainers/base:0-ubuntu-18.04

# ** [Optional] Uncomment this section to install additional packages. **
#
# ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
   && apt-get -y install --no-install-recommends golang fish \
   #
   # Clean up
   && apt-get autoremove -y \
   && apt-get clean -y \
   && rm -rf /var/lib/apt/lists/*
# ENV DEBIAN_FRONTEND=dialog
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

