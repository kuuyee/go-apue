package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
)

func main() {
	fileInfo, err := os.Stat("foo")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%b\n", fileInfo.Mode())!
	//fmt.Printf("%b\n", fileInfo.Mode()&unix.S_IXGRP)
	err = os.Chmod("foo", fileInfo.Mode()|unix.S_ISGID|unix.S_IXGRP)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Chmod("bar", unix.S_IRUSR|unix.S_IWUSR|unix.S_IRGRP|unix.S_IROTH)
	if err != nil {
		fmt.Println(err)
	}
}
