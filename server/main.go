package main

import (
	"EK-Server/app"
	"EK-Server/config"
	"EK-Server/model"
	"EK-Server/util"
	"fmt"
	"os"

	"github.com/fvbock/endless"
)

func main() {
	e := app.New()

	//初始化数据链接
	if err := model.MysqlConn(config.Global.Mysql); err != nil {
		panic(err)
	}
	defer model.DB.Close() //退ß出时释放链接

	// e.Logger.Error(e.Start(fmt.Sprintf(":%d", config.Global.Port)))

	server := endless.NewServer(fmt.Sprintf(":%d", config.Global.Port), e)
	err := server.ListenAndServe()
	if err != nil {
		util.Logger.Println(err)
	}
	fmt.Printf("stopd\n")
	os.Exit(0)
}
