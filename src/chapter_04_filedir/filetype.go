package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Args Number must be > 2")
		os.Exit(1)
	}

	for _, v := range os.Args {
		fileInfo, err := os.Stat(v)
		//fileInfo, err := os.Lstat(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Name : ", fileInfo.Name())

		fmt.Println("Size : ", fileInfo.Size())

		fmt.Println("Mode/permission : ", fileInfo.Mode())

		//for os.Lstat
		if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
			fmt.Println("File is a symbolic link")
		}

		fmt.Println("Modification Time : ", fileInfo.ModTime())

		fmt.Println("Is a directory? : ", fileInfo.IsDir())

		fmt.Println("Is a regular file? : ", fileInfo.Mode().IsRegular())

		fmt.Println("Unix permission bits? : ", fileInfo.Mode().Perm())

		fmt.Println("Permission in string : ", fileInfo.Mode().String())

		fmt.Println("What else underneath? : ", fileInfo.Sys())
		fmt.Println()
		fmt.Println()
	}
}
