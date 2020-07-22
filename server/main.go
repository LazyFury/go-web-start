package main

import (
	"EK-Server/app"
	"EK-Server/config"
	"EK-Server/model"
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

	server := endless.NewServer(fmt.Sprintf(":%d", config.Global.Port), e)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("stopd\n")
	// config.Global.Mail.SendMail("服务异常断开", []string{"suke971219@gmail.com"}, fmt.Sprintf("%+v", os.Getpid()))
	os.Exit(0)
}
