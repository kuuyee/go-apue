package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("参数不能少于2个！")
		os.Exit(1)
	}
	fmt.Println(os.Args[1])
	unix.Access(os.Args[1], unix.W_OK)
	if false {
		fmt.Println("access error for ", os.Args[1])
	} else {
		fmt.Println("read access OK ")
	}
}
