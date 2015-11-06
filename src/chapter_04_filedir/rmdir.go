package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("tempdir")
	if err != nil {
		fmt.Errorf("删除目录报错：%s", err)
	}

	err = os.RemoveAll("abc")
	if err != nil {
		fmt.Errorf("删除目录报错：%s", err)
	}
}
