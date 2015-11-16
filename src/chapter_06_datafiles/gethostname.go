package main

import (
	"fmt"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Errorf("获取系统hostname报错：%s", err)
		return
	}
	fmt.Printf("系统Hostname: %s\n", hostname)
}
