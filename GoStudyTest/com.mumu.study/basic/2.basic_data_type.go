// @Author liuzhen
// @Date 2023/12/16 17:41:00
// @Desc
package main

import "fmt"

/**
1. 整型
	1.1. 按长度分为：int8、int16、int32、int64
	1.2. 对应的无符号整型：uint8、uint16、uint32、uint64
	其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型。

2. 浮点型
	两种浮点型数：float32和float64。
	这两种浮点型数据格式遵循IEEE 754 标准：
		float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。
		float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64。

3. 复数
	3.1. complex64
	3.2. complex128

4. 布尔值
	布尔型数据只有true（真）和false（假）两个值。
	注意：
		布尔类型变量的默认值为false。
		Go 语言中不允许将整型强制转换为布尔型.
		布尔型无法参与数值运算，也无法与其他类型进行转换。

5. 字符串
	5.1. 定义单行字符串使用双引号("")。定义多行字符串使用 `反引号`。
   	5.2. 字符串转义符:
				 \r	回车符（返回行首）
				 \n	换行符（直接跳到下一行的同列位置）
				 \t	制表符
				 \'	单引号
				 \"	双引号
				 \\	反斜杠
	5.3.

6. byte和rune类型
	6.1. uint8类型，或者叫 byte 型，代表一个ASCII码字符。
		 rune类型，代表一个 UTF-8字符。
	6.2. 修改字符串: 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。

7. 类型转换。
   强制类型转换的基本语法如下: T(表达式)


*/

func intDemo() {
	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) // 把十进制数转换成二进制
	fmt.Printf("%o\n", i1) // 把十进制数转换成八进制
	fmt.Printf("%x\n", i1) // 把十进制数转换成十六进制

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)
	// 查看变量的类型
	fmt.Printf("%T\n", i3)

	// 声明int8类型的变量
	i4 := int8(9) // 明确指定int8类型，否则就是默认为int类型
	fmt.Printf("%T\n", i4)
}

func floatDemo() {
	// math.MaxFloat32 // float32最大值
	f1 := 1.23456
	fmt.Printf("%T\n", f1) // 默认Go语言中的小数都是float64类型
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2) // 显示声明float32类型
	// f1 = f2                // float32类型的值不能直接复赋值给float64类型的变量
}

func boolDemo() {
	b1 := true
	var b2 bool // 默认是false
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T value:%v\n", b2, b2)
}

func main() {

}
