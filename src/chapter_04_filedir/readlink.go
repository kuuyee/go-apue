package main

import (
	"fmt"
	"os"
)

func main() {
	linkInfo, err := os.Readlink("foolink")
	if err != nil {
		fmt.Errorf("读取链接报错：", err)
		return
	}
	fmt.Println(linkInfo)
}
