// @Author liuzhen
// @Date 2023/12/15 16:18:00
// @Desc
package main

import "fmt"

/**
1. 变量声明
	1.1. 标准声明 格式: var 变量名 变量类型
	2.2. 批量声明 格式: var (变量名 变量类型
						  变量名 变量类型
					    )

2. 变量的初始化
	2.1. 标准格式: var 变量名 类型 = 表达式
	2.2. 类型推导 格式: var 变量名 = 值
	3.3. 短变量声明: 在函数内部，可以使用更简略的 := 方式声明并初始化变量。格式: 变量名 := 值

3. 匿名变量(anonymous variable): 名变量用一个下划线_表示 eg: _ = 1234

4. 常量: 把var换成了const，常量在定义的时候必须赋值
	1.1. 单个: const pi = 3.1415
	2.2. 多个常量也可以一起声明：
		const (
    		pi = 3.1415
    		e = 2.7182
		)

5. iota: 是go语言的常量计数器，只能在常量的表达式中使用。
	  iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
	TODO 介绍: https://www.cnblogs.com/zsy/p/5370052.html

6. 内置函数介绍
	6.1. close(c chan<- Type)  主要用来关闭channel
	6.2. len(v Type) int	  用来求长度，比如: string、array、slice、map、channel
		 cap(v Type) int
	6.3. new(Type) *Type	  用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
	6.4. make(t Type, size ...IntegerType) Type	  用来分配内存，主要用来分配引用类型，比如chan、map、slice
	6.5. append(slice []Type, elems ...Type) []Type 用来追加元素到数组、slice中
		 copy(dst, src []Type) int
		 delete(m map[Type]Type1, key Type)
		 print(args ...Type)
		 println(args ...Type)
		 complex(r, i FloatType) ComplexType
		 real(c ComplexType) FloatType
		 imag(c ComplexType) FloatType

	6.6. panic(v interface{}) 和 recover() interface{}	用来做错误处理。
		Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。 panic可以在任何地方引发，但 recover只有在defer调用的函数中有效。
		注意：
			recover() 必须搭配defer使用。
			defer 一定要在可能引发 panic 的语句之前定义。

*/

// 1. 变量 ============================>
// Go语言中推荐使用驼峰式命名
// var student_name string
var studentName string

// 批量声明
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

// 2. 常量 ============================>
// 定义了常量之后不能修改
// 在程序运行期间不会改变的量
const pi = 3.1415926

// 批量声明常量
const (
	statusOK = 200
	notFound = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

// 3. iota ============================>
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)

const (
	b1 = iota // 0
	b2 = iota // 1
	_  = iota // 2
	b3 = iota // 3
)

// 插队
const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota // 2
	c4
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // d1:1 d2:2
	d3, d4 = iota + 1, iota + 2 // d3:2 d4:3
)

// 定义数量级
// 这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	// 2.1. 声明变量同时赋值
	var s1 string = "whb"
	fmt.Println(s1)
	// 2.2. 类型推导（根据值判断该变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)
	// 2.3. 简短变量声明，只能在函数里面用
	s3 := "哈哈哈"
	fmt.Println(s3)

	// pi = 123
	// fmt.Println("n1:", n1)
	// fmt.Println("n2:", n2)
	// fmt.Println("n3:", n3)

	// fmt.Println("a1:", a1)
	// fmt.Println("a2:", a2)
	// fmt.Println("a3:", a3)

	// fmt.Println("b1:", b1)
	// fmt.Println("b2:", b2)
	// fmt.Println("b3:", b3)

	// fmt.Println("c1:", c1)
	// fmt.Println("c2:", c2)
	// fmt.Println("c3:", c3)
	// fmt.Println("c4:", c4)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)

	// panic demo
	panicFuncA()
	panicFuncB()
	panicFuncC()
}

func panicFuncA() {
	fmt.Println("func A")
}

func panicFuncB() {
	defer func() {
		err := recover()

		fmt.Println(err)
		fmt.Println("释放数据库连接...")

		// 如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()

	// 程序崩溃退出
	panic("panic in B")

	fmt.Println("func B")
}

func panicFuncC() {
	fmt.Println("func C")
}
