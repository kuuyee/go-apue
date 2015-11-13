package main

import (
	"fmt"
	"time"
)

func main() {
	localtime := time.Now()
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	fmt.Printf("当前时间为：%s\n", localtime.Format(layout))
	fmt.Printf("当前时间(Unix格式): %s\n", localtime.Format(time.UnixDate))
}
