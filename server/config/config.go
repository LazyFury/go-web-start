package config

import (
	"EK-Server/util"
	"EK-Server/util/wechat"
	"encoding/json"
	"log"
	"os"
)

// Global 全局配置
var Global *configType = initConfig()

func initConfig() *configType {
	t := &configType{}
	//读取配置文件
	if err := t.ReadConfig(); err != nil {
		panic(err)
	}
	return t
}

type configType struct {
	// 数据库链接
	Mysql string `json:"mysql"`
	// 网站根目录
	BaseURL     string    `json:"baseURL"`
	TablePrefix string    `json:"tablePrefix"`
	Wechat      wechat.MP `json:"wechat"`
	Port        int       `json:"port"`
	Mail        util.Mail `json:"mail"`
}

// ReadConfig 读取配置 初始化时运行 绑定为全局变量
// 在我使用 ReadConfig 命名函数的时候 编辑器提示了错误， 函数应该和结构体configType保存一直的大写或者小写 以保证其他包的调用者可以使用这个函数
func (c *configType) ReadConfig() (err error) {
	f, err := os.Open("./config/config.json")
	defer f.Close()
	if err != nil {
		log.Fatalln("打开配置文件错误，请创建 config/config.json 参考(config-defaultjson")
		return
	}
	if err = json.NewDecoder(f).Decode(c); err != nil {
		return
	}
	return nil
}
