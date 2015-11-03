## 文件和目录

### access.c
在"golang.org/x/sys/unix"包中找到相关的函数定义，但是只有err一个返回值，不知道怎么判断access函数的设置效果。

```go
func Access(path string, mode uint32) (err error) {
	return Faccessat(AT_FDCWD, path, mode, 0)
}
```

需要注意的是，这个包没法在windows下调试，sublime里没法点取函数，并且无法编译，必须在unix环境下才有效


### umask(文件模式创建屏蔽符)
umask的概念可以参考：[http://linux.vbird.org/linux_basic/0220filemanager.php#umask](http://linux.vbird.org/linux_basic/0220filemanager.php#umask)

```go
package main

import (
	"fmt"
	"golang.org/x/sys/unix" //必须在Unix环境执行
	"os"
)

func main() {
	unix.Umask(0)
	_, err := os.Create("foo")
	if err != nil {
		fmt.Println("Create Error")
	}
	unix.Umask(unix.S_IRGRP | unix.S_IWGRP | unix.S_IROTH | unix.S_IWOTH)
	_, err2 := os.Create("bar")

	if err2 != nil {
		fmt.Println("Create Error")
	}
}
```

### changemod
没有找到c程序`(statbuf.st_mode & ~S_IXGRP) | S_ISGID)`对应的golang写法，所以简化为`fileInfo.Mode()|unix.S_ISGID|unix.S_IXGRP`，反正不影响理解chmod函数。

```go
package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
)

func main() {
	fileInfo, err := os.Stat("foo")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%b\n", fileInfo.Mode())!
	//fmt.Printf("%b\n", fileInfo.Mode()&unix.S_IXGRP)
	err = os.Chmod("foo", fileInfo.Mode()|unix.S_ISGID|unix.S_IXGRP)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Chmod("bar", unix.S_IRUSR|unix.S_IWUSR|unix.S_IRGRP|unix.S_IROTH)
	if err != nil {
		fmt.Println(err)
	}
}
```

验证

```
[vagrant@mydev ~]$ ls -l foo bar
-rw-r--r--. 1 vagrant vagrant 0 Oct 26 02:54 bar
-rw-rw-rw-. 1 vagrant vagrant 0 Oct 26 02:54 foo
[vagrant@mydev ~]$ 
[vagrant@mydev ~]$ go run changemod.go 
[vagrant@mydev ~]$ 
[vagrant@mydev ~]$ ls -l foo bar       
-rw-r--r--. 1 vagrant vagrant 0 Oct 26 02:54 bar
-rw-rwxrw-. 1 vagrant vagrant 0 Oct 26 02:54 foo //可以看到foo的权限设置成功了
```

### unlink
golang标准库中没有unlink，因此需要使用`golang.org/x/sys/unix`包

```go
package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
)

func main() {
	_, err := os.Open("tempfile")
	if err != nil {
		fmt.Println(err)
	}
	err = unix.Unlink("tempfile")
	if err != nil {
		fmt.Println(err)
	}
}
```

验证

```
[vagrant@mydev ~]$ ls -l tempfile 
-rw-rw-r--. 1 vagrant vagrant 38677 Oct 28 02:26 tempfile
[vagrant@mydev ~]$ 
[vagrant@mydev ~]$ df /home/
Filesystem              1K-blocks    Used Available Use% Mounted on
/dev/mapper/centos-root  14571520 4780728   9790792  33% /
[vagrant@mydev ~]$ 
[vagrant@mydev ~]$ go run unlink.go 
[vagrant@mydev ~]$ ls -l tempfile
ls: cannot access tempfile: No such file or directory	//unlink成功，无法找到文件
[vagrant@mydev ~]$ 
[vagrant@mydev ~]$ ls
bar  changemod.go  ci  docker  foo  go  go1.4.2.linux-amd64.tar.gz  go-apue  local  umask.go  unlink  unlink.go
[vagrant@mydev ~]$ 
[vagrant@mydev ~]$ df /home/
Filesystem              1K-blocks    Used Available Use% Mounted on
/dev/mapper/centos-root  14571520 4780692   9790828  33% /	//磁盘占用量少了
```

### rename
golang提供的对应的函数[os.Rename](https://gowalker.org/os#Rename)

```go
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	oldFile, err := os.Create("bar")
	if err != nil {
		fmt.Fprintf(os.Stdout, "创建文件报错：", err)
	}
	oldFile.Close()
	time.Sleep(time.Second * 5)

	//改文件
	err = os.Rename("bar", "foo")
	if err != nil {
		fmt.Fprintf(os.Stdout, "创建文件报错：", err)
	}

	//改目录
	err = os.Rename("oldpath", "newpath")
	if err != nil {
		fmt.Fprintf(os.Stdout, "创建文件报错：", err)
	}
}
```

> 如果new文件已经存在则报错：`创建文件报错：%!(EXTRA *os.LinkError=rename bar foo: Cannot create a file when that file already exists.)`

### symlink
golang提供的对应的函数[os.Symlink](https://gowalker.org/os#Symlink)

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Symlink("foo", "foolink")
	if err != nil {
		fmt.Fprintf(os.Stdout, "创建文件连接报错：", err)
	}

	fmt.Println(os.Readlink("foolink"))
}
```

### readlink
如果读取的不是软连接则返回空

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	linkInfo, err := os.Readlink("foolink")
	if err != nil {
		fmt.Errorf("读取链接报错：", err)
		return
	}
	fmt.Println(linkInfo)
}
```

### futimens
golang没有futimens函数，但是有Chtimes：[https://gowalker.org/os#Chtimes](https://gowalker.org/os#Chtimes)

```go
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
```

验证时间戳更改

```
$ ls -la
-rwxrwxrwx. 1 vagrant vagrant    0 Nov  3 01:49 foo // 时间戳 3 01:49

$ go run futimens.go 
foo最后修改时间是：2015-11-03 01:49:58 +0000 WET
清空后的文件修改时间是：2015-11-03 01:57:53 +0000 WET // 时间戳改变
[vagrant@mydev chapter_04_filedir]$ ls -l
-rwxrwxrwx. 1 vagrant vagrant    0 Nov  3 01:49 foo	 // 时间戳被还原了
```




### ftw8
golang提供了遍历目录的函数walk:[https://gowalker.org/path/filepath#Walk](https://gowalker.org/path/filepath#Walk)

```go
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
```

测试：

```
$ go run ftw8.go /home/vagrant/go-apue/src
Root参数：  /home/vagrant/go-apue/src
/home/vagrant/go-apue/src
/home/vagrant/go-apue/src/chapter_03_fileio
/home/vagrant/go-apue/src/chapter_03_fileio/1.txt
/home/vagrant/go-apue/src/chapter_03_fileio/README.md
/home/vagrant/go-apue/src/chapter_03_fileio/fileflags.go
/home/vagrant/go-apue/src/chapter_03_fileio/hole.c
/home/vagrant/go-apue/src/chapter_03_fileio/hole.go
/home/vagrant/go-apue/src/chapter_03_fileio/mycat.go
/home/vagrant/go-apue/src/chapter_03_fileio/seek.c
/home/vagrant/go-apue/src/chapter_03_fileio/seek.go
/home/vagrant/go-apue/src/chapter_04_filedir
/home/vagrant/go-apue/src/chapter_04_filedir/README.md
/home/vagrant/go-apue/src/chapter_04_filedir/a.out
/home/vagrant/go-apue/src/chapter_04_filedir/access.go
/home/vagrant/go-apue/src/chapter_04_filedir/bar
/home/vagrant/go-apue/src/chapter_04_filedir/changemod.go
/home/vagrant/go-apue/src/chapter_04_filedir/filetype.go
/home/vagrant/go-apue/src/chapter_04_filedir/foo
/home/vagrant/go-apue/src/chapter_04_filedir/foolink
/home/vagrant/go-apue/src/chapter_04_filedir/ftw8.go
/home/vagrant/go-apue/src/chapter_04_filedir/futimens.go
/home/vagrant/go-apue/src/chapter_04_filedir/newpath
/home/vagrant/go-apue/src/chapter_04_filedir/readlink.go
/home/vagrant/go-apue/src/chapter_04_filedir/rename.go
/home/vagrant/go-apue/src/chapter_04_filedir/symlink.go
/home/vagrant/go-apue/src/chapter_04_filedir/u01
/home/vagrant/go-apue/src/chapter_04_filedir/u02
/home/vagrant/go-apue/src/chapter_04_filedir/umask.go
/home/vagrant/go-apue/src/chapter_04_filedir/unlink.go
-----------------------
目录总数：5
文件总数：23
```