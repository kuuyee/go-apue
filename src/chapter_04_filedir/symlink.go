package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Symlink("foo", "foolink")
	if err != nil {
		fmt.Fprintf(os.Stdout, "创建文件报错：", err)
	}

	fmt.Println(os.Readlink("foolink"))
}
