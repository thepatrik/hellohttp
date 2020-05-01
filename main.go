package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/thepatrik/hellohttp/pkg/server"
)

func main() {
	_ = godotenv.Load()
	cfg := newConf()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	svr, err := server.New(
		server.WithLogger(logger),
		server.WithPort(cfg.Port),
	)
	if err != nil {
		logger.Fatal(err)
	}

	done := make(chan bool)
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	go func() {
		sig := <-gracefulStop
		logger.Printf("Caught signal (%s). Gracefully shutting down...\n", sig)

		if err := svr.Stop(); err != nil {
			logger.Printf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	logger.Printf("Running HTTP server on %s", svr.Addr())
	if err := svr.Start(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not start server on %s: %v\n", svr.Addr(), err)
	}
	<-done
	logger.Print("Server stopped")
}
