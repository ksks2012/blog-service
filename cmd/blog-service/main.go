package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/blog-service/pkg/tracer"

	"github.com/blog-service/global"
	"github.com/blog-service/internal/dao/config"
	"github.com/blog-service/internal/model"
	"github.com/blog-service/internal/routers"
	"github.com/blog-service/pkg/logger"
	"github.com/blog-service/pkg/setting"
	"github.com/gin-gonic/gin"
	"golang.org/x/sys/unix"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	cfg string
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = updateDB()
	if err != nil {
		log.Fatalf("init.updateDB err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}

}

// @title 部落格系統
// @version 1.0
// @description Go 語言編程之旅：一起用 Go 做項目
// @termsOfService https://github.com
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, os.Interrupt, unix.SIGTERM)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

func setupSetting() error {
	s, err := setting.NewSetting(strings.Split(cfg, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second

	return nil
}

func updateDB() error {
	var err error
	updateDBSetup := &config.StorageSetup{}
	err = updateDBSetup.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	if err = updateDBSetup.Instance.Open(); nil != err {
		log.Fatalf("open storage connection failed: %v", err)
		return err
	}

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
