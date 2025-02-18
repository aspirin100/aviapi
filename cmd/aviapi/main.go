package main

import (
	"fmt"
	"log"

	"github.com/aspirin100/aviapi/internal/config"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("current config:", cfg)

	// application, err := app.New(context.Background(), cfg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = application.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// stop := make(chan os.Signal, 1)
	// signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// <-stop

	// err = application.Stop(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("server correctly stopped")
}
