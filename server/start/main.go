package main

import (
	"fmt"
	"os"

	"github.com/Treblex/go-echo-demo/server/app"
	"github.com/Treblex/go-echo-demo/server/config"
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/Treblex/go-echo-demo/server/router"
	"github.com/Treblex/go-echo-demo/server/util/mlog"

	"github.com/fvbock/endless"
)

func main() {
	e := app.New()

	//初始化数据链接
	if err := model.MysqlConn(config.Global.Mysql.ToString()); err != nil {
		panic(err)
	}
	defer model.DB.Close() //退出时释放链接

	// 静态目录
	e.Static("/", "wwwroot")

	// 注册路由
	router.Start(e)

	// e.Logger.Error(e.StartTLS(fmt.Sprintf(":%d", config.Global.Port), "cert.pem", "key.pem"))
	server := endless.NewServer(fmt.Sprintf(":%d", config.Global.Port), e)
	err := server.ListenAndServe()
	if err != nil {
		mlog.Error(err.Error())
	}
	mlog.Error("stopd\n")
	os.Exit(0)
}
