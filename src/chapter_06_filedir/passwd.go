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
	pw_name   string //用户名
	pw_passwd string //加密口令
	pw_uid    string //用户ID
	pw_gid    string //用户组ID
	pw_gecos  string //注释字段
	pw_dir    string //初始工作目录
	pw_shell  string //初始Shell程序
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
