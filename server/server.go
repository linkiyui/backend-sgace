package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	api_server "github.com/sgace/api_server"
)

func Start() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		s := <-sigc
		fmt.Println("terminating ATcnea server | signal: ", s)
		os.Exit(0)
	}()

	api_server.NewApiServer().Start()
}
