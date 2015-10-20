package main

import (
	"log"
	"os"
)

func main() {
	holefile, err := os.OpenFile(
		"file.hole",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE, //如果文件不存在就新建一个
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer holefile.Close()

	_, err = holefile.WriteString("abcdefghij")
	if err != nil {
		log.Fatal(err)
	}
	_, err = holefile.Seek(16385, 0) //设置偏移，生成一个空洞
	if err != nil {
		log.Fatal(err)
	}
	_, err = holefile.WriteString("ABCDEFGHIJ")
	if err != nil {
		log.Fatal(err)
	}
}
