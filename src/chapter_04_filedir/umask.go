package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
)

func main() {
	unix.Umask(0)
	_, err := os.Create("foo")
	if err != nil {
		fmt.Println("Create Error")
	}
	unix.Umask(unix.S_IRGRP | unix.S_IWGRP | unix.S_IROTH | unix.S_IWOTH)
	_, err2 := os.Create("bar")

	if err2 != nil {
		fmt.Println("Create Error")
	}
}
