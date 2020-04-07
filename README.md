# rays-gin-framework
簡介:rays-gin-framework 整合了 gorm 以及 gin 的基礎網頁框架可以快速的產出一個小型的web app 
包含了一些web server基礎的功能: sessions,config配置,graceful shutdown。搭配cli可以快速建構prototype，之後會陸續增加一些常用的middleware
以及component

專案結構:
```
./rays-gin-framework
├── LICENSE
├── README.md
├── config
│   └── config.go
├── config.ini
├── database
│   └── database.go
├── go.mod
├── go.sum
├── main.go
├── pkg
│   └── util
│       └── utiliy.go
├── router
│   └── router.go
└── templates
    └── index.html
``` 

所有的網頁Router都從router.go進入，可依照需求寫成mvc或是單純的web function
```
func RegisterRouter(engine *gin.Engine) {
	//TODO Register App Router or Register Api Router
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
```

database 的用法也寫成全域變數放便在其他地方呼叫
```
func TestFunc(){
   var user model.User
   database.Context.Find(&user)
   fmt.Println(user)
}
```

templates資料夾中放置網頁模版，因gin的模板載入方式所以資料夾階層需為單層
```
templates
    ├── index.html
    ├── index1.html
    ├── index2.html
    ├── index3.html
    └── index4.html
```

