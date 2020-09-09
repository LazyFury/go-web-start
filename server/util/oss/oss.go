package oss

import (
	"fmt"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// AliyunOssUpload AliyunOssUpload
func AliyunOssUpload(name string, file io.Reader) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(fmt.Sprintf("Err:%x", err))
		}
	}()

	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "s36RGPs6Tw378idZ", "mqrzXqwASefnnfLQEZKocmhyYplA9I")
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

	return "https://suke100.oss-cn-beijing.aliyuncs.com/" + name
}
