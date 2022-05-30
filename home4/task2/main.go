package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handle(signal os.Signal) {
	fmt.Println("Received:", signal)
	if signal == syscall.SIGTERM {
		time.Sleep(2 * time.Second)
	}
}

func main() {
	sigs := make(chan os.Signal, 1)
	fmt.Println(syscall.SIGTERM)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				handle(sig)
			case syscall.SIGTERM:
				handle(sig)
			case syscall.SIGINT:
				handle(sig)
				os.Exit(0)
			default:
				fmt.Println("Ignoring:", sig)
			}
		}
	}()
	for {
		fmt.Println("Doing a stupid job")
		time.Sleep(3 * time.Second)
	}

	//fmt.Printf("received %s\n", s)
}
