package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Chdir("newpath")
	if err != nil {
		fmt.Errorf("改变工作目录错误：%s", err)
	}
	fmt.Println("当前工作目录更改为：newpath")
	err = os.Mkdir("tempdir", 0666)
	if err != nil {
		fmt.Errorf("创建目录报错：%s", err)
	}
}
