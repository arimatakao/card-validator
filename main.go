package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/arimatakao/card-validator/server"
)

func main() {
	srv := server.New(":8080")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := srv.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Error occurred while running server: ", err.Error())
		} else {
			log.Println("Shutdown server")
		}
	}()

	log.Println("Server started and listeting on port :8080")

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown with error: ", err.Error())
	}

	log.Println("Shutdown is successful")
	os.Exit(0)
}
