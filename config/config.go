package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Global 全局配置
var Global *configType = readConfig()

// configType ConfigType
type configType struct {
	// 数据库链接
	Mysql string `json:"mysql"`
	// 网站根目录
	BaseURL string `json:"baseURL"`
}

// ReadConfig 读取配置 初始化时运行 绑定为全局变量
// 在我使用 ReadConfig 命名函数的时候 编辑器提示了错误， 函数应该和结构体保存一直的大写或者小写 以保证其他包的调用者可以使用这个函数
func readConfig() *configType {
	f, err := os.Open("./config/config.json")
	defer f.Close()
	if err != nil {
		panic(fmt.Errorf("打开文件错误"))
	}
	c := configType{}
	if err = json.NewDecoder(f).Decode(&c); err != nil {
		panic(err)
	}
	return &c
}
