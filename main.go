package main

import (
	"github.com/alexperezortuno/toolbox/cmd"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := make(chan os.Signal)
	signal.Notify(ctx, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-ctx
		os.Exit(0)
	}()

	cmd.Execute()
}
