package main

import (
	"log"
	"os"
	"os/signal"

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

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "ping"})
	// })
	// r.Run(":18080")
}
