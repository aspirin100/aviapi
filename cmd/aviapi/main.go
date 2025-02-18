package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aspirin100/aviapi/internal/app"
	"github.com/aspirin100/aviapi/internal/config"
)

const (
	shutdownTimeout = time.Second * 5
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("current config:", cfg)

	application, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	stopCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	go func() {
		err = application.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	err = application.Stop(stopCtx)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("correctly stopped")
}
