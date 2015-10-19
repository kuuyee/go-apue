package main

import (
	"fmt"
	"os"
)

func main() {
	b := make([]byte, 10)
	f, err := os.Open("1.txt")
	defer f.Close()
	if err == nil {
		offset, err := f.Seek(1, 0) //偏移1位
		if err == nil {
			n, err := f.Read(b)
			if err == nil {
				fmt.Printf("offset: %d\n读取偏移的内容：%s\n", offset, string(b[:n]))
			}
		}
	}
}
