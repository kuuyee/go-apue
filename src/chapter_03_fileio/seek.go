package main

import (
	"fmt"
	"os"
)

//Seek设置下一次读/写的位置。offset为相对偏移量，而whence决定相对位置：
//0为相对文件开头，1为相对当前位置，2为相对文件结尾。它返回新的偏移量（相对开头）和可能的错误。
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
