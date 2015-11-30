## 进程环境

### 7.3 atexit进程终止

> 需要在Unix环境下执行才能看到输出

```go
    package main

    import (
        "fmt"
        "os"
        "os/signal"
    )

    func main() {
        signalChan := make(chan os.Signal, 1)
        cleanupDone := make(chan bool)
        signal.Notify(signalChan, os.Interrupt)
        go func() {
            for _ = range signalChan {
                fmt.Println("\n收到终端信号，停止服务... \n")
                cleanup()
                cleanupDone <- true
            }
        }()
        <-cleanupDone
    }

    func cleanup() {
        fmt.Println("清理...\n")
    }
```

执行程序，并`Ctrl+c`看效果

    $ go run atexit3.go
    ^C
    收到终端信号，停止服务...

    清理...

### 7.11 函数getrlimit和setrlimit
标准包中没有对应函数，所以需要使用`syscall`

```go
  package main

  import (
  	"fmt"
  	"os"
  	"syscall"
  )

  func main() {
  	var rlim syscall.Rlimit
  	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlim)
  	if err != nil {
  		fmt.Errorf("获取Rlimit报错 %s", err)
  		os.Exit(1)
  	}
  	fmt.Printf("ENV RLIMIT_NOFILE : %+v\n", rlim)

  	rlim.Max = 65535
  	rlim.Cur = 65535
  	//err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlim)
  	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlim)
  	if err != nil {
  		fmt.Errorf("设置Rlimit报错 %s", err)
  		os.Exit(1)
  	}
  	//fmt.Printf("ENV RLIMIT_NOFILE : %+v\n", rlim)
  	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlim)
  	if err != nil {
  		fmt.Errorf("获取Rlimit报错 %s", err)
  		os.Exit(1)
  	}
  	fmt.Printf("ENV RLIMIT_NOFILE : %+v\n", rlim)
  }
```

执行输出

    # go run getrlimit.go
    ENV RLIMIT_NOFILE : {Cur:1024 Max:4096}
    ENV RLIMIT_NOFILE : {Cur:65535 Max:65535}
