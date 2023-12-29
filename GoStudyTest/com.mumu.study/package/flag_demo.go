// @Author liuzhen
// @Date 2023/12/29 15:14:00
// @Desc
package main

import (
	"flag"
	"fmt"
	"time"
)

/**
flag包: 实现命令行参数的解析

1. flag参数类型: flag包支持的命令行参数类型有bool、int、int64、uint、uint64、float float64、string、duration。
	flag参数		有效值
	字符串flag		合法字符串
	整数flag		1234、0664、0x1234等类型，也可以是负数。
	浮点数flag		合法浮点数
	bool类型flag	1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False。
	时间段flag		任何合法的时间段字符串。如"300ms"、"-1.5h"、“2h45m”。合法的单位有"ns"、“us” /“µs”、“ms”、“s”、“m”、“h”。

2. 常用api:
	2.1. flag.Type()  定义命令行参数方式
	基本格式如下: flag.Type(flag名, 默认值, 帮助信息) *Type
		eg: 定义姓名、年龄、婚否三个命令行参数
		需要注意的是，此时name、age、married、delay均为对应类型的指针。

	2.2. flag.TypeVar()  定义命令行参数方式
	基本格式如下: flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)

	2.3. flag.Parse(): 解析命令行参数
		通过以上两种方法定义好命令行flag参数后，需要通过调用 flag.Parse() 来对命令行参数进行解析。
		支持的命令行参数格式有以下几种:
			-flag xxx （使用空格，一个-符号）
			--flag xxx （使用空格，两个-符号）
			-flag=xxx （使用等号，一个-符号）
			--flag=xxx （使用等号，两个-符号）
		其中，布尔类型的参数必须使用等号的方式指定。
		Flag解析在第一个非flag参数（单个"-“不是flag参数）之前停止，或者在终止符”–“之后停止。

	2.4. flag.Args()  返回命令行参数后的其他参数，以[]string类型
	2.5. flag.NArg()  返回命令行参数后的其他参数个数
	2.6. flag.NFlag() 返回使用的命令行参数个数

3. 使用
	3.1. 命令行参数使用提示: go run flag_demo.go -help
	3.2. 正常使用命令行flag参数: go run flag_demo.go -name 沙河娜扎 --age 28 -married=false -d=1h30m
	3.3. 使用非flag命令行参数: go run flag_demo.go a b c		获取参数: 2.4 flag.Args()

*/

func main() {
	// 2.1 flag.Type()  定义姓名、年龄、婚否三个命令行参数，方式1
	//name0 := flag.String("name", "张三", "姓名")
	//age0 := flag.Int("age", 18, "年龄")
	//married0 := flag.Bool("married", false, "婚否")
	//delay0 := flag.Duration("d", 0, "时间间隔")
	//fmt.Printf("name0: %v age0: %v married0: %v delay0: %v\n", name0, age0, married0, delay0)

	// 2.2. flag.TypeVar()  定义姓名、年龄、婚否三个命令行参数，方式2
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	// 2.3. flag.Parse() 解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)

	// 2.4. 返回命令行参数后的其他参数
	fmt.Println(flag.Args())

	// 2.5. 返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	// 2.6. 返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
