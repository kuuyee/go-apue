package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("tempdir", 0666)
	if err != nil {
		fmt.Errorf("创建目录报错：%s", err)
	}

	err = os.MkdirAll("abc/def/g", 0666)
	if err != nil {
		fmt.Errorf("创建目录报错：%s", err)
	}
}
