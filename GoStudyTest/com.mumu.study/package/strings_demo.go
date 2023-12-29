// @Author liuzhen
// @Date 2023/12/19 17:38:00
// @Desc
package main

import (
	"fmt"
	"strings"
)

/**
字符串的常用操作
	len(str)	求长度
	+ 或 fmt.Sprintf()	拼接字符串
	strings.Split()	分割
	strings.contains()	判断是否包含
	strings.HasPrefix(), strings.HasSuffix()	前缀/后缀判断
	strings.Index(), strings.LastIndex()	子串出现的位置
	strings.Join(a[]string, sep string)	join操作
*/

func main() {
	s3 := `D:\Go\src\code.oldboyedu.com\studygo\day01`

	// 字符串相关操作
	fmt.Println(len(s3)) // ?

	// 字符串拼接
	name := "理想"
	world := "大帅比"
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	// fmt.Printf("%s%s", name, world)
	fmt.Println(ss1)

	// 分隔
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	// 包含
	fmt.Println(strings.Contains(ss, "理性"))
	fmt.Println(strings.Contains(ss, "理想"))

	// 前缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	// 后缀
	fmt.Println(strings.HasSuffix(ss, "理想"))

	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))

	// 拼接
	fmt.Println(strings.Join(ret, "+"))
}
