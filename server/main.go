package main

import (
	"fmt"
	"os"

	"github.com/treblex/go-echo-demo/server/app"
	"github.com/treblex/go-echo-demo/server/config"
	"github.com/treblex/go-echo-demo/server/model"
	"github.com/treblex/go-echo-demo/server/util"

	"github.com/fvbock/endless"
)

func main() {
	e := app.New()
	//初始化数据链接
	if err := model.MysqlConn(config.Global.Mysql.ToString()); err != nil {
		panic(err)
	}
	defer model.DB.Close() //退出时释放链接
	// e.Logger.Error(e.StartTLS(fmt.Sprintf(":%d", config.Global.Port), "cert.pem", "key.pem"))
	server := endless.NewServer(fmt.Sprintf(":%d", config.Global.Port), e)
	err := server.ListenAndServe()
	if err != nil {
		util.Logger.Println(err)
	}
	fmt.Printf("stopd\n")
	os.Exit(0)
}
