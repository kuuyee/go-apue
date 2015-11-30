package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			fmt.Println("\n收到终端信号，停止服务... \n")
			cleanup()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}

func cleanup() {
	fmt.Println("清理...\n")
}
