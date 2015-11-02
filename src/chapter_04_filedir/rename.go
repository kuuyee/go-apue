package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	oldFile, err := os.Create("bar")
	if err != nil {
		fmt.Fprintf(os.Stdout, "创建文件报错：", err)
	}
	oldFile.Close()
	time.Sleep(time.Second * 5)

	//改文件
	err = os.Rename("bar", "foo")
	if err != nil {
		fmt.Fprintf(os.Stdout, "创建文件报错：", err)
	}

	//改目录
	err = os.Rename("oldpath", "newpath")
	if err != nil {
		fmt.Fprintf(os.Stdout, "创建文件报错：", err)
	}
}
