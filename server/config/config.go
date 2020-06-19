package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Global 全局配置
var Global *configType = readConfig()

type (
	wechat struct {
		Appid     string `json:"appid"`
		Appsecret string `json:"appsecret"`
	}
	configType struct {
		// 数据库链接
		Mysql string `json:"mysql"`
		// 网站根目录
		BaseURL     string `json:"baseURL"`
		TablePrefix string `json:"tablePrefix"`
		Wechat      wechat `json:"wechat"`
		Port        int    `json:"port"`
		Mail        mail
	}
)

// ReadConfig 读取配置 初始化时运行 绑定为全局变量
// 在我使用 ReadConfig 命名函数的时候 编辑器提示了错误， 函数应该和结构体configType保存一直的大写或者小写 以保证其他包的调用者可以使用这个函数
func readConfig() *configType {
	f, err := os.Open("./config/config.json")
	defer f.Close()
	if err != nil {

		defaultConf := &configType{
			Port:  8080,
			Mysql: "[username]:[password]@(localhost:3306)/[databaseName]?charset=utf8mb4&parseTime=true&loc=Asia%2fShanghai",
		}
		b, _ := json.Marshal(defaultConf)
		fmt.Println(string(b))

		f, err = os.Create("./config/config.json")
		if err != nil {
			panic(err)
		}

		f.Write(b)

		panic("打开配置文件错误，请补充填写 config/config.json 参考(config-defaultjson")
	}
	c := configType{}
	if err = json.NewDecoder(f).Decode(&c); err != nil {
		panic(err)
	}
	return &c
}
