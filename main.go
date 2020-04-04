
package main

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"rays-gin-framework/config"
	"rays-gin-framework/database"
	"rays-gin-framework/router"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//設定log format
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})

	//讀入config.ini 配置檔案
	setting := config.NewConfig("config.ini")

	//Database連線
	database.NewDatabase(setting.Database)
	defer database.CloseDataBase()

	//設定Gin Mode
	gin.SetMode(setting.Runtime.Mode)

	//初始化 cookie base Session
	store := cookie.NewStore([]byte("dvY6BxVP"))

	//初始化 Gin Router Engine
	engine := gin.New()

	//如果./templates沒有檔案 49行會出錯
	engine.LoadHTMLGlob("./templates/*.html")

	engine.Use(gin.Logger(), gin.Recovery())
	engine.Use(sessions.Sessions("rays-gin-framework", store))

	//註冊網頁路由
	router.RegisterRouter(engine)

	//Graceful shutdown server
	server := &http.Server{
		Addr:    setting.Runtime.Port,
		Handler: engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quitSignal := make(chan os.Signal,1)
	signal.Notify(quitSignal,syscall.SIGINT,syscall.SIGTERM)
	<- quitSignal

	log.Println("Shutdown Server ...")

	ctx , cancel := context.WithTimeout(context.Background(), 5 *time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}

