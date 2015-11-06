package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	dev := []interface{}{os.ModeDir, os.ModeAppend, os.ModeExclusive, os.ModeTemporary, os.ModeSymlink, os.ModeDevice, os.ModeNamedPipe, os.ModeSocket, os.ModeSetuid, os.ModeSetgid, os.ModeCharDevice, os.ModeSticky, os.ModeType, os.ModePerm}
	for _, v := range dev {
		devType := reflect.TypeOf(v)
		devVal := reflect.ValueOf(v)
		fmt.Printf("设备类型为：%v | 值：%v\n", devType, devVal.Interface())
	}
}
