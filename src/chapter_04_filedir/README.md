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

