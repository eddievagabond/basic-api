package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	server "github.com/eddievagabond/internal"
	"github.com/eddievagabond/internal/storage"
	"github.com/eddievagabond/internal/util"
)

func main() {
	c := util.NewConfiguration()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	stopper := make(chan struct{})
	go func() {
		<-done
		close(stopper)
	}()

	store, err := storage.NewStorage(c)
	if err != nil {
		log.Fatalf("Error creating storage: %s", err)
	}

	apiServer, err := server.NewApiServer(c.ServerAddress, store)
	if err != nil {
		log.Fatalf("error creating server: %s", err)
	}
	apiServer.Start(stopper)
}
