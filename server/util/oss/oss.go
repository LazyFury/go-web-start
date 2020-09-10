package oss

import (
	"fmt"
	"io"

	"github.com/Treblex/go-echo-demo/server/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// AliyunOssUpload AliyunOssUpload
func AliyunOssUpload(name string, file io.Reader) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(fmt.Sprintf("Err:%x", err))
		}
	}()
	conf := config.Global.AliOss
	client, err := oss.New(conf.Endpoint, conf.AccessKeyID, conf.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	bucket, err := client.Bucket("suke100")
	if err != nil {
		panic(err)
	}
	err = bucket.PutObject(name, file)
	if err != nil {
		panic(err)
	}

	return conf.URL + name
}
