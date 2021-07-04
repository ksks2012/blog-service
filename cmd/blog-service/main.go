package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/blog-service/global"
	"github.com/blog-service/internal/routers"
	"github.com/blog-service/pkg/setting"
	"golang.org/x/sys/unix"
)

var (
	config string
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

}

func main() {
	// cfg, err := parseCommandParam()
	// if nil != err {
	// 	log.Fatalf("invalid command option for blog-service: %v", err)
	// 	return
	// }
	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, os.Interrupt, unix.SIGTERM)
	log.Printf("INFO: starting Blog service")
	// if err = cfg.StorageSetup.Instance.Open(); nil != err {
	// 	log.Fatalf("open storage connection failed: %v", err)
	// 	return
	// }

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
	s, err := setting.NewSetting(strings.Split(config, ",")...)
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

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}
