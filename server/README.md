# Go-echo-demo
Go语言 学习 demo

## go语言入门
推荐这个 https://tour.golang.org/welcome/2 还没看完 嘿嘿
Go语言在线练习场，讲解很细

##  web框架
简单看了一些基础我就来搞这个了，这个项目也是一个web实践

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

#  📃关于这个项目

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

## 🦊杂项 遇到的一些小问题

### 时间戳格式化
```go
    var timeLayout string = "2006年01月02日 15:04:05"
    t := time.Now()
    t.Format(timeLayout)

    //这里 2006 01 02 15 04 05 必须是固定的才能解析到正确的时间
    //我在 util/time.go 内实现了一个简单的方法以使用 y-m-d h:i:s 来格式化时间 主要还是因为懒得记
    //  Format 也提供了很多的默认layout，默认layou英文支持做的好一些，中文的显示自定义的layout满足需求
```

### 数据库存取时间类型

#### gorm 时间类型自定义解析格式

```go 

// LocalTime 继承time.Time类型
type LocalTime struct {
	time.Time
}
// 自定义的 layout
var timeLayout string = "2006年01月02日 15:04:05"

// MarshalJSON json格式化时间的方法
// 在网上搜到的方式是格式化为时间戳的  不符合我的需求，这个方法是从 time源代码里找到的,直接修改默认 layout为自定义即可
func (t LocalTime) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	b := make([]byte, 0, len(timeLayout)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, timeLayout)
	b = append(b, '"')
	return b, nil
}


//下边两个方法是gorm需要,不需要修改，如果没有使用 gorm则不需要
// Value Value
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan Scan
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
```

####   存取时间的时候相差8小时
```go
//DataBase 数据库配置 username:password@host/database_name?param
DataBase string = "root:2568597007suke@(localhost:3306)/test?charset=utf8mb4&parseTime=true&loc=Asia%2fShanghai"

//charset=utf8mb4 数据库编码
//parseTime=true    自动解析时间 time.Time 类型解析显示不正确  z100:h1223 之类的一个字符串
//loc=Asia%2fShanghai  默认亚洲时间，数据库存储 detatime 默认为utc时区 也就是会比国内早8个小时
```

####    查询数据时 隐藏某些隐私字段
gorm:"-" 在保存数据的时候会忽略 查询是同样显示

```go
    //暂时的解决方案是声明两个模型，在查询时仅显示必要字段
    //由于gorm使用结构名 + s  例如 users articles 默认为约定表名，所以声明另外的模型是需要在查询数据只指定表名
    // gorm.Table("users")
```


####  ！gorm 更新 或者 添加删除字段后 查询影响的数据行数一直为0，因此无法知道是否更新成功 或者 保存数据成功

在操作之后重新赋值db 以获取新的位置   而不是直接使用 db.RowsAffected
```go
row := db.Model(&User{ID: id}).Updates(data)
row.RowsAffected
```

###  go get 安装完包之后不能使用命令后直接执行
path 设置错误
```sh
# //go mod需要开启
export GO111MODULE=on 
# //指定代理
export GOPROXY=https://goproxy.cn; 
# goroot  go环境所在目录  系统包
export GOROOT=/usr/local/go	
# 安装的package所在目录  自己定义，修改目录之后需要重新安装package
export GOPATH=~/gowork
# 将GOPATH/bin 加入全局path之后 自定义安装的package就可以全局运行了 fresh govender
export PATH=$GOPATH/bin:$GOPATH:$PATH

```

###	web开发时自动重启项目
[https://github.com/gravityblast/fresh](https://github.com/gravityblast/fresh)
看起来像是检查ctrl s 的时候重新编译文件并执行，文件没有修改的时候也会触发重启

##  ⚠️暂未解决的问题

## 环境设置 
```
export GO111MODULE=on GOPROXY=https://goproxy.cn;
export GOROOT=/usr/local/go
export GOPATH=~/gowork

export PATH=$GOPATH/bin:$GOPATH:$PATH

```