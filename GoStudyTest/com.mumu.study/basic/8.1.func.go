// @Author liuzhen
// @Date 2023/12/19 20:23:00
// @Desc
package main

import "fmt"

/**
函数

1. 函数定义: 使用func关键字，具体格式如下：
	func 函数名(参数) (返回值) {
    	函数体
	}
	其中:
		函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名。
		参数：参数由参数变量和参数变量的类型组成，多个参数之间使用,分隔。
		返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔。
		函数体：实现指定功能的代码块。

2. 函数调用: 直接通过 函数名() 的方式调用函数

3. 参数:
	3.1. 类型简写: 函数的参数中如果相邻变量的类型相同，则可以省略类型。eg: func intSum(x, y int) {}
	3.2. 可变参数: 可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。注意：可变参数通常要作为函数的最后一个参数。
		eg: func intSum2(x ...int) {}

4. 返回值: 通过return关键字向外输出返回值。
	4.1. 多返回值: Go语言中函数支持多返回值，函数如果有多个返回值时必须用()将所有返回值包裹起来。eg: func calc(x, y int) (int, int) {...}
	4.2. 返回值命名: 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。eg: func calc(x, y int) (sum, sub int) {...}
	4.3. 返回值补充: 当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。
		eg: func someFunc(x string) []int {
				if x == "" {
					return nil // 没必要返回[]int{}
				}
				...
			}

5. 变量作用域
	5.1. 全局变量: 全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量。eg: 在go文件里定义的变量，var = num int64 = 100
	5.2. 局部变量: 函数内 或 语句块定义的变量(if条件判断、for循环、switch语句上使用这种定义变量)

6. 函数类型与变量
	6.1. 定义函数类型: 使用type关键字来定义一个函数类型，具体格式如下: type calculation func(int, int) int
		上面语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。
		简单来说，凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型。

	6.2. 函数类型变量: 我们可以声明函数类型的变量并且为该变量赋值
		eg: 声明一个calculation类型的变量c: var c calculation
			把add赋值给c: c = add

7. 高阶函数
	7.1. 函数作为参数
	7.2. 函数作为返回值

8. 匿名函数和闭包
	8.1. 匿名函数: 匿名函数就是没有函数名的函数
		格式如下: func(参数) (返回值) {
    				函数体
				  }
	8.2. 匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数。匿名函数多用于实现回调函数和闭包。
		eg: noNameFunction()

9. 闭包
	闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包 = 函数 + 引用环境。
	闭包是什么？
		闭包是一个函数，这个函数包含了他外部作用域的一个变量
	底层的原理：
		1. 函数可以作为返回值
		2. 函数内部查找变量的顺序，先在自己内部找，找不到往外层找

10. defer语句
	将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
	defer执行时机: 它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。


*/

// 定义一个全局变量
var globalVal = 100

type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 定义一个函数
func f11() {
	// globalVal := 10
	name := "理想"
	// 函数中查找变量的顺序
	// 1. 先在函数内部查找
	// 2. 找不到就往函数的外面查找,一直找到全局
	fmt.Println(globalVal, name)
}

// 定义一个函数:
// 1. 当参数中连续多个参数的类型一致时，我们可以将非最后一个参数的类型省略
// 2. 可变长参数: 可变长参数必须放在函数参数的最后
// 3. 多个带命名的返回值
func f(x, y, z int, k ...float64) (a int, b string) {
	return x + y, ""

}

func main() {
	f11()
	// fmt.Println(name) // 函数内部定义的变脸只能在该函数内部使用

	// 语句块作用域
	if i := 10; i < 18 {
		fmt.Println("乖乖上学")
	}
	// fmt.Println(i) // 不存在i
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
	// fmt.Println(j) // 不存在j

	// 闭包demo
	var f = adderBibao()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	f1 := adderBibao()
	fmt.Println(f1(40)) //40
	fmt.Println(f1(50)) //90

	//
	var c calculation               // 声明一个calculation类型的变量c
	c = add                         // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c

	f2 := add                        // 将函数add赋值给变量f
	fmt.Printf("type of f:%T\n", f2) // type of f:func(int, int) int
	fmt.Println(f2(10, 20))          // 像调用add一样调用f
}

// 匿名函数变量
var noNameVal = func(x, y int) {
	fmt.Println(x + y)
}

// 匿名函数demo
func noNameFunction() {
	// 1. 将匿名函数保存到变量。函数内部没有办法声明带名字的函数
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	// 2. 自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

}

// 闭包函数
func adderBibao() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
