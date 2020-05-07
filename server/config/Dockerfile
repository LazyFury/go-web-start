FROM alpine

WORKDIR /root/app
COPY  zoneinfo.zip /opt/zoneinfo.zip
ENV ZONEINFO /opt/zoneinfo.zip
EXPOSE 8080
CMD ["./linux-main"]