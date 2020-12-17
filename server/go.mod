module github.com/Treblex/go-echo-demo/server

go 1.13

require (
	github.com/Treblex/go-web-template v0.0.0-20201217075835-7b0cbbec72d1
	github.com/Treblex/simple-daily v0.0.0-20201214014444-2874a70d28a6
	github.com/aliyun/aliyun-oss-go-sdk v2.1.5+incompatible
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/static v0.0.0-20200916080430-d45d9a37d28e
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.1.2
	github.com/gorilla/websocket v1.4.2
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/ugorji/go v1.2.0 // indirect
	golang.org/x/crypto v0.0.0-20201124201722-c8d3bf9c5392 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.8
)

replace github.com/Treblex/go-web-template => ../../go-web-template
