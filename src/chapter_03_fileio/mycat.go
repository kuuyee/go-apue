package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入，并以回车结束")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Errorf("Read input Error!")
		return
	}
	fmt.Print("您输入的内容是：")
	os.Stdout.WriteString(input)
}
