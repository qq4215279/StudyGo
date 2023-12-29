// @Author liuzhen
// @Date 2023/12/20 11:42:00
// @Desc
package main

import (
	"fmt"
	"os"
)

/**
os包
1. os.Args 获取命令行参数



*/

// os.Args demo
// 将下面的代码执行: go run os_demo.go a b c
func main() {
	//os.Args是一个[]string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}
