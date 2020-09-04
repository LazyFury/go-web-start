package upload

import (
	"github.com/treblex/go-echo-demo/server/util"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

// Default 默认上传类型和文件夹
func Default(c echo.Context) error {
	return Custom(c, []string{}, "")
}

// Custom 自定义上传类型和目录
func Custom(c echo.Context, acceptsExt []string, folder string) error {
	file, err := c.FormFile("file")
	if err != nil {
		return util.JSONErr(c, err, "上传错误") //未获取到文件流
	}
	link, err := uploadBase(file, acceptsExt, folder)
	if err != nil {
		fmt.Println(err)
		errMsg := fmt.Sprintf("%s", err)
		if errMsg == "" {
			errMsg = "上传失败"
		}
		return util.JSONErr(c, nil, errMsg)
	}
	return util.JSONSuccess(c, link, "上传成功")
}

// acceptsExt  这里是一个扩展到类型，默认到图片，视频 压缩包类型，已经写在默认方法中了
func uploadBase(file *multipart.FileHeader, acceptsExt []string, folderName string) (fileName string, err error) {
	pathExt := path.Ext(file.Filename)

	folder := ""
	// 如果符合类型，设定目录
	if inArray(AcceptsImgExt, strings.Trim(pathExt, ".")) {
		folder = "image"
	}
	if inArray(AcceptsVideoExt, strings.Trim(pathExt, ".")) {
		folder = "video"
	}
	if inArray(AcceptsAudioExt, strings.Trim(pathExt, ".")) {
		folder = "audio"
	}
	if inArray(AcceptsOtherFileExt, strings.Trim(pathExt, ".")) {
		folder = "file"
	}
	// 自定义类型  覆盖前边的
	if inArray(acceptsExt, strings.Trim(pathExt, ".")) {
		folder = folderName
	}
	// 如果不符合任何一种类型
	if folder == "" {
		err = errors.New("文件不合法")
		return
	}

	// 打开文件流
	src, err := file.Open()
	if err != nil {
		err = errors.New("打开文件失败")
		return

	}
	defer src.Close() //函数结束时自动关闭文件

	//创建文件夹
	dir, err := util.GetDir("./static/upload/"+folder+"/", time.Now().Format("2006_01_02"))
	if err != nil {
		err = errors.New("创建文件夹失败")
		return
	}

	// 随机文件名 + 文件后缀
	randName := util.RandStringBytes(32) + pathExt
	// Destination
	fileName = filepath.Join(dir, randName)

	// 创建空文件
	dst, err := os.Create(fileName)
	if err != nil {
		err = errors.New("创建文件失败")
		return
	}
	defer dst.Close()
	// Copy文件流到新建到文件
	if _, err = io.Copy(dst, src); err != nil {
		err = errors.New("拷贝文件至目标失败")
		return
	}
	// 拼接文件地址，不带协议头，方便处理http 到https升级 ， 其实也没找到协议头在哪儿，req对象里没有返回到空字符串
	fileName = fmt.Sprintf("/%s", fileName)
	return
}

// 在数组中
func inArray(arr []string, item string) (inArr bool) {
	index := -1
	item = strings.ToLower(item)
	for i, x := range arr {
		if item == x {
			index = i
		}
	}
	return index > -1
}
