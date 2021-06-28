package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/blog-service/internal/routers"
	"golang.org/x/sys/unix"
)

func main() {
	cfg, err := parseCommandParam()
	if nil != err {
		log.Fatalf("invalid command option for blog-service: %v", err)
		return
	}
	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, os.Interrupt, unix.SIGTERM)
	log.Printf("INFO: starting Blog service")
	if err = cfg.StorageSetup.Instance.Open(); nil != err {
		log.Fatalf("open storage connection failed: %v", err)
		return
	}

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":18080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
