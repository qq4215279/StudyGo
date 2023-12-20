// @Author liuzhen
// @Date 2023/12/19 17:09:00
// @Desc
package main

import "fmt"

/**
1. 向外输出
	1.1. Print 函数会将内容输出到系统的标准输出
		Print() 函数直接输出内容: func Print(a ...interface{}) (n int, err error)
		Printf() 函数支持格式化输出字符串: func Printf(format string, a ...interface{}) (n int, err error)
		Println() 函数会在输出内容的结尾添加一个换行符: func Println(a ...interface{}) (n int, err error)

	1.2. Fprint Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
		func Fprint(w io.Writer, a ...interface{}) (n int, err error)
		func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
		func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

	1.3. Sprint Sprint系列函数会把传入的数据生成并返回一个字符串。
		func Sprint(a ...interface{}) string
		func Sprintf(format string, a ...interface{}) string
		func Sprintln(a ...interface{}) string

	1.4. Errorf Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
		func Errorf(format string, a ...interface{}) error

2. 获取输入 运行过程中从标准输入获取用户的输入。
	2.1 Scan
		func Scan(a ...interface{}) (n int, err error)
		func Scanf(format string, a ...interface{}) (n int, err error)
		func Scanln(a ...interface{}) (n int, err error)

	2.2. Fscan 这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从io.Reader中读取数据。
		func Fscan(r io.Reader, a ...interface{}) (n int, err error)
		func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
		func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)

	2.3. Sscan系列 这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从指定字符串中读取数据。
		func Sscan(str string, a ...interface{}) (n int, err error)
		func Sscanln(str string, a ...interface{}) (n int, err error)
		func Sscanf(str string, format string, a ...interface{}) (n int, err error)

3. 格式化占位符  *printf系列函数都支持format格式化参数
	3.1. 通用占位符:
		%v	值的默认格式表示
		%+v	类似%v，但输出结构体时会添加字段名
		%#v	值的Go语法表示
		%T	打印值的类型
		%%	百分号

	3.2. 布尔型
		%t	true 或 false

	3.3. 整型
		%b	表示为二进制
		%c	该值对应的unicode码值
		%d	表示为十进制
		%o	表示为八进制
		%x	表示为十六进制，使用a-f
		%X	表示为十六进制，使用A-F
		%U	表示为Unicode格式：U+1234，等价于"U+%04X"
		%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示

	3.4. 浮点数与复数
		%b	无小数部分、二进制指数的科学计数法，如-123456p-78
		%e	科学计数法，如-1234.456e+78
		%E	科学计数法，如-1234.456E+78
		%f	有小数部分但无指数部分，如123.456
		%F	等价于%f
		%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
		%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）

	3.5. 字符串和[]byte
		%s	直接输出字符串或者[]byte
		%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
		%x	每个字节用两字符十六进制数表示（使用a-f
		%X	每个字节用两字符十六进制数表示（使用A-F）

	3.6. 指针
		%p	表示为十六进制，并加上前导的0x

	3.7. 宽度标识符
	宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。精度通过（可选的）宽度后跟点号后跟的十进制数指定。
	如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0。举例如下：
		%f	默认宽度，默认精度
		%9f	宽度9，默认精度
		%.2f	默认宽度，精度2
		%9.2f	宽度9，精度2
		%9.f	宽度9，精度0

	3.8. 其他flag
		‘+’	总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
		’ '	对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格
		‘-’	在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
		‘#’	八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值；
		‘0’	使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；

*/

// fmt占位符
func main() {
	var n = 100
	// 查看类型
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)

	var s = "Hello 沙河！"
	fmt.Printf("字符串：%s\n", s)
	fmt.Printf("字符串：%v\n", s)
	fmt.Printf("字符串：%#v\n", s)
}

func fmtDemo() {
	// fmt.Print("沙河")
	// fmt.Print("娜扎")
	// fmt.Println("--------")
	// fmt.Println("沙河")
	// fmt.Println("娜扎")
	// Printf("格式化字符串", 值)
	// %T :查看类型
	// %d :十进制数
	// %b :二进制数
	// %o :八进制数
	// %x :十六进制数
	// %c :字符
	// %s :字符串
	// %p： 指针
	// %v： 值
	// %f：浮点数
	// %t ：布尔值

	// var m1 = make(map[string]int, 1)
	// m1["理想"] = 100
	// fmt.Printf("%v\n", m1)
	// fmt.Printf("%#v\n", m1)

	// printBaifenbi(90)

	// fmt.Printf("%v\n", 100)
	// // 整数->字符
	// fmt.Printf("%q\n", 65)
	// // 浮点数和复数
	// fmt.Printf("%b\n", 3.14159265354697)
	// // 字符串
	// fmt.Printf("%q\n", "李想有理想")
	// fmt.Printf("%7.3s\n", "李想有理想")

	// 获取用户输入
	var s string
	fmt.Scan(&s)
	fmt.Println("用户输入的内容是：", s)

	// var (
	// 	name  string
	// 	age   int
	// 	class string
	// )
	// // fmt.Scanf("%s %d %s\n", &name, &age, &class)
	// // fmt.Println(name, age, class)

	// fmt.Scanln(&name, &age, &class)
	// fmt.Println(name, age, class)

	//fmt.Printf("%b\n", 1024)

}

func printBaifenbi(num int) {
	fmt.Printf("%d%%\n", num)
}
