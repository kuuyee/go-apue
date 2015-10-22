## FileIO

### fcntl实现
go语言并没有c语言中fcntl函数的对应实现，参看golang官方的说明:[https://github.com/golang/go/issues/487](https://github.com/golang/go/issues/487)

因此需要借用go中的底层调用包syscall来实现，代码如下：

```go
package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "请使用：fileflag  0 < /dev/tty ", os.Args[0])
	}
	fmt.Println(os.Args[1])
	fd, err := strconv.Atoi(os.Args[1])
	if err != nil {
		val, err := Fcntl(fd, 3, 0)
		if err != nil {
			fmt.Printf("val : ", val)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func Fcntl(fd, cmd int, arg int) (val int, err error) {
	r, _, e := syscall.Syscall(syscall.SYS_FCNTL, uintptr(fd), uintptr(cmd), uintptr(arg))
	val = int(r)
	if e != 0 {
		err = e
	}
	return
}
```

编译虽然通过，但是没有返回我预期的结果，还有待进一步研究。