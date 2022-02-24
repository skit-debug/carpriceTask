package main

import (
	"context"
	"main/src/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctx, cancelServer := context.WithCancel(ctx)
	defer cancelServer()
	go server.StartServer(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	go func() {
		<-c
		cancelServer()
	}()
}
