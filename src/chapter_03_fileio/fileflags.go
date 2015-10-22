package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "请使用：fileflag  0 < /dev/tty ", os.Args[0])
	}
	fmt.Println(os.Args[1])
	fd, err := strconv.Atoi(os.Args[1])
	if err != nil {
		val, err := Fcntl(fd, 3, 0)
		if err != nil {
			fmt.Printf("val : ", val)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func Fcntl(fd, cmd int, arg int) (val int, err error) {
	r, _, e := syscall.Syscall(syscall.SYS_FCNTL, uintptr(fd), uintptr(cmd), uintptr(arg))
	val = int(r)
	if e != 0 {
		err = e
	}
	return
}
