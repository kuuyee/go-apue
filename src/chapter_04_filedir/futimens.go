package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fooInfo, err := os.Stat("foo")
	if err != nil {
		fmt.Errorf("读取文件状态报错 %s", err)
		return
	}
	initTime := fooInfo.ModTime()
	fmt.Printf("foo最后修改时间是：%v\n", initTime)

	fooFile, err := os.OpenFile("foo", os.O_TRUNC, 0666)
	if err != nil {
		fmt.Errorf("打开文件报错 %s", err)
		return
	}
	ftime, _ := fooFile.Stat()
	fmt.Printf("清空后的文件修改时间是：%v\n", ftime.ModTime())
	fooFile.Close()

	err = os.Chtimes("foo", time.Now(), initTime)
	if err != nil {
		fmt.Errorf("修改时间戳报错 %s", err)
	}
}
