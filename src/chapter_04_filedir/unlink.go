package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
)

func main() {
	_, err := os.Open("tempfile")
	if err != nil {
		fmt.Println(err)
	}
	err = unix.Unlink("tempfile")
	if err != nil {
		fmt.Println(err)
	}
}
