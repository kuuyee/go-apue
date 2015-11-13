## 系统数据文件和信息

### 6.2 口令文件
golang标准库中没有`getpwuid`的实现，只有`func Getuid() int`,不能返回整体的Passwd结构，下面用go实现了`/etc/passwd`文件的解析

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Passwd结构包含/etc/passwd的七项内容
type Passwd struct {
	pw_name   string	//用户名
	pw_passwd string	//加密口令
	pw_uid    string	//用户ID
	pw_gid    string	//用户组ID
	pw_gecos  string	//注释字段
	pw_dir    string	//初始工作目录
	pw_shell  string	//初始Shell程序
}

func main() {
	pwfile, err := os.Open("/etc/passwd")
	if err != nil {
		fmt.Errorf("读取/etc/passwd报错：", err)
	}
	defer pwfile.Close()

	pwf, err := ParsePasswdFile(pwfile)
	fmt.Printf("root信息：%+v\n", pwf["root"])

}

func ParsePasswdFile(r io.Reader) (map[string]Passwd, error) {
	pwline := bufio.NewReader(r)
	pwMap := make(map[string]Passwd)
	for {
		line, _, err := pwline.ReadLine()
		if err != nil {
			break
		}
		pwArray := strings.Split(string(line), ":")
		if len(pwArray) != 7 {
			fmt.Errorf("读取用户passwd信息报错：")
			return nil, err
		}
		passwd := new(Passwd)
		passwd.pw_name = pwArray[0]
		passwd.pw_passwd = pwArray[1]
		passwd.pw_uid = pwArray[2]
		passwd.pw_gid = pwArray[3]
		passwd.pw_gecos = pwArray[4]
		passwd.pw_dir = pwArray[5]
		passwd.pw_shell = pwArray[6]
		pwMap[passwd.pw_name] = *passwd

	}
	return pwMap, nil
}

```

验证，获取root用户的passwd信息

```
$ go run passwd.go 
root信息：{pw_name:root pw_passwd:x pw_uid:0 pw_gid:0 pw_gecos:root pw_dir:/root pw_shell:/bin/bash}
```


### 6.9 gethostname

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Errorf("获取系统hostname报错：%s", err)
		return
	}
	fmt.Printf("系统Hostname: %s\n", hostname)
}
```

### 6.10 strftime

```go
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
```

输出

```
$ go run strftime.go 
当前时间为：Nov 13, 2015 at 1:33am (WET)
当前时间(Unix格式): Fri Nov 13 01:33:12 WET 2015
```