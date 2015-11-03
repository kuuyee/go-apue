package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var nreg, ndir int

type myfunc filepath.WalkFunc

func main() {
	flag.Parse()
	root := flag.Arg(0)
	fmt.Println("Root参数： ", root)

	err := filepath.Walk(root, filepath.WalkFunc(myftw()))
	if err != nil {
		fmt.Errorf("遍历目录报错: %s", err)
	}
	fmt.Println("-----------------------")
	fmt.Printf("目录总数：%d\n", ndir)
	fmt.Printf("文件总数：%d\n", nreg)

}

func myftw() myfunc {
	return func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		fmt.Println(path)
		infoMode := info.Mode()

		if infoMode.IsRegular() {
			nreg++
		}
		if infoMode.IsDir() {
			ndir++
		}
		return nil
	}
}
