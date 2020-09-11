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
	// 相对目录
	path = fmt.Sprintf("/%s", fileName)
	return
}

func defaultGetFile(req *http.Request) (header *multipart.FileHeader, err error) {
	_, header, err = req.FormFile("file")
	return
}
