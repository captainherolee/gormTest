package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	router "gormTest/router"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, os.Interrupt, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGKILL)

	go router.StartServer()

	sig := <-c
	log.Printf("Got %s signal. Aborting...\n", sig)
}
