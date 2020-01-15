# go-echo-demo
go语言 学习 demo

## go语言入门
推荐这个 https://tour.golang.org/welcome/2 还没看完 嘿嘿
go语言在线练习场，讲解很细

##  web框架
简单看了写基础我就来搞这个了，这个项目也是一个web实践

>   ### echo   [官方文档](http://go-echo.org/) 注意首页有几个很有用的简单实例，但是你按照文档目录是无法返回这一页的 😂 当时找了好久找不到
>   这个项目就是    据说超快

> ### gin   [官方文档](https://gin-gonic.com/zh-cn/docs/)
>   江湖人称：用过都说好，据说开发是可以自动重启 不用手动运行项目 想试试，但是我懒  和上边那个一样快，又一个扩展项目 好像叫green 可以自动生成文档，不过很久没更新了，有兴趣了解一下这个怎么实现的

> ###   iris    [官方文档](https://iris-go.com/)
>   大型项目支持

##  数据库支持
[gorm中文文档](https://jasperxu.github.io/gorm-zh/crud.html#q) 我看的这个 因为搜索引擎先找到了这个  后来才发现其实不是官方的 这里>>
[gorm官方文档](https://gorm.io/) 
```go
        "github.com/jinzhu/gorm" //数据库操作框架 支持模型和链式操作 不用写sql了
        _ "github.com/go-sql-driver/mysql" //go语言sql驱动
        
        db, err := gorm.Open("mysql", config.DataBase)
        if err != nil {
            panic(err)
        }
        db.LogMode(true)
        defer db.Close() //defer表示函数结束是调用 ，  但是我不想在每个函数都重新打开链接，所以我创建了一个 util包 并暴露了 DB 供全局使用，我在 main.go 的函数结尾出调用了 defer db.Close()，但我并不知道他是否关闭了链接，可以知道的是链接确实只创建了一次，并一直保持
    
```

#  关于这个项目

```go
.
├── Dockerfile //docker构建 我还在摸索 不过确实运行成功了
├── README.md   
├── build.linux.sh  //linux构建文件，我在一段时间后才意识到 编译到不同平台到二进制文件也会有不同，不该是0和1吗 😂
├── build.sh         //mac构建
├── build.win.sh    //win .exe文件
├── config          
│   └── config.go   //配置文件
├── dist            //编译后到文件
│   └── main-linux  
├── go.mod          //go模块
├── go.sum          //git clone 之后可能需要删除这个文件重新按照mod
├── h5              // 这是一个简单到请求接口到页面，在尝试了一段时间layui写admin页面之后，觉得直接写接口好了，>>> mvc 主要是卡在 模版分离之后 header 和 footer 之类到layout文件 如何请求数据，因为他不对应一个路由，当前找到到方案是模版函数 但是没有运行成功 可能我绑定模版render到方法也有问题, 也尝试了ifarme 不太喜欢，还是后边试一下antd或者自己搞一下vue或者react的后台页面
│   ├── api.js
│   ├── index.html
│   ├── main.css
│   └── main.js
├── main.go         //入口文件
├── modal           //模型文件， 数据库查询的方法
│   └── user.go
├── router          //路由文件
│   ├── admin       //内包含一个类似router.go的入口文件，和其他路由文件
│   └── router.go   //暴露一个 Init 方法接收 echo对象, 用于声明当前路面的子文件夹路由对象 或者 路由 
>>>>>> 例子 >>>>>>>>>
    // 项目首页
	admin.Init(e, baseURL)

	// 入口
	e.GET(baseURL, func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world！")
    })
>>>>>>>>>>>>>>>
├── static  //资源文件  
├── template    //模板文件 弃用mvc之后 只保留了error
│   └── error.html
└── util    //工具包
    ├── db.go   // 暴露一个DB对象 用户操作数据库,启动后仅链接一次
    ├── return.go   //返回的json对象格式 和错误码声明
    ├── template.go //声明模版
    ├── time.go //时间格式化  以及  链接数据库是自动格式化的一个 结构体
    └── util.go //暂无内容--
```