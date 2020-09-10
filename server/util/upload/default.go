package upload

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func defaultUpload(fileName string, src io.Reader) (path string, err error) {
	// 创建空文件
	dst, err := os.Create(fileName)
	if err != nil {
		err = errors.New("创建文件失败")
		return
	}
	defer dst.Close()
	// Copy文件流到新建到文件
	if _, err := io.Copy(dst, src); err != nil {
		err = errors.New("拷贝文件至目标失败")
	}
	// 拼接文件地址，不带协议头，方便处理http 到https升级 ， 其实也没找到协议头在哪儿，req对象里没有返回到空字符串
	path = fmt.Sprintf("/%s", fileName)
	return
}

func defaultGetFile(req *http.Request) (header *multipart.FileHeader, err error) {
	_, header, err = req.FormFile("file")
	return
}
