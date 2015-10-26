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