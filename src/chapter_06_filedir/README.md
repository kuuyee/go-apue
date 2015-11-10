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
	pw_name   string
	pw_passwd string
	pw_uid    string
	pw_gid    string
	pw_gecos  string
	pw_dir    string
	pw_shell  string
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