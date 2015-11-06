package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Chdir("dirlink")
	if err != nil {
		fmt.Errorf("改变工作目录错误：%s", err)
	}

	workdir, _ := os.Getwd()
	fmt.Printf("当前工作目录为： %s\n", workdir)

}
