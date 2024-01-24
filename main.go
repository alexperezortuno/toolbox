package main

import (
	"fmt"
	"github.com/alexperezortuno/toolbox/cmd"
	"os"
	"os/signal"
	"syscall"
)

var version = "0.0.1"

func main() {
	fmt.Printf("Toolbox: v%s\n", version)
	ctx := make(chan os.Signal)
	signal.Notify(ctx, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-ctx
		os.Exit(0)
	}()

	cmd.Execute()
}
