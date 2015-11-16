package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var exitFuncs []func()

// TrapSignals注册一个处理器要来捕获信号和终止信号
// 在os.Exit(1)之后调用退出函数
func TrapSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		CallExitFuncs()
		os.Exit(1)
	}()
}

func Run(f func()) {
	exitFuncs = append(exitFuncs, f)
}

func CallExitFuncs() {
	fmt.Println("Comein CallExitFuncs")
	for i := len(exitFuncs) - 1; i >= 0; i-- {
		exitFuncs[i]()
	}
}

func main() {
	TrapSignals()
	defer CallExitFuncs()
	Run(func() {
		fmt.Println("1 atexit!")
	})
	Run(func() {
		fmt.Println("2 atexit!")
	})

}
