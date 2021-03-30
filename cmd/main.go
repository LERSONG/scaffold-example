package main

import (
	"github.com\LERSONG\scaffold-example/config"
	"github.com\LERSONG\scaffold-example/model"
	"github.com\LERSONG\scaffold-example/web"

	"log"
	"os"
	"os/signal"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to reading config file, %s\n", err)
	}

	service, err := model.New(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize model for operating all service, %s\n", err)
	}

	server := web.NewServer(cfg, service)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen for http server, %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	log.Println("scaffold-example is running")
	<-quit
	log.Println("scaffold-example is stopped")
}
