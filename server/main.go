package main

import (
	"fmt"

	"github.com/Treblex/go-echo-demo/server/app"
	"github.com/Treblex/go-echo-demo/server/config"
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/Treblex/go-echo-demo/server/router"
)

func main() {
	e := app.New()

	//初始化数据链接
	if err := model.MysqlConn(config.Global.Mysql.ToString()); err != nil {
		panic(err)
	}
	defer model.DB.Close() //退出时释放连接
	// 静态目录
	e.Static("/", "wwwroot")

	// 注册路由
	router.Start(e)

	e.Logger.Error(e.Start(fmt.Sprintf(":%d", config.Global.Port)))

}
