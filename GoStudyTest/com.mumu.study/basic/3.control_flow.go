// @Author liuzhen
// @Date 2023/12/19 17:48:00
// @Desc
package main

import "fmt"

/**
控制语句

1. if 条件
	if 表达式1 {
		分支1
	} else if 表达式2 {
		分支2
	} else {
		分支3
	}

2. for循环
	2.1. 基本格式如下：
		for 初始语句; 条件表达式; 结束语句 {
			循环体语句
		}

	2.2. goto(跳转到指定标签)
		goto 语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。
		Go语言中使用goto语句能简化一些代码的实现过程。 例如双层嵌套的for循环要退出时

	2.3. break(跳出循环)
		break语句可以结束for、switch和select的代码块。

	2.4. continue(继续下次循环)
		continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用。

3. 无限循环 for循环可以通过break、goto、return、panic语句强制退出循环。
	for {
		循环体语句
	}

4. for range(键值循环)
	Go语言中可以使用for range遍历数组、切片、字符串、map 及通道（channel）。 通过for range遍历的返回值有以下规律：
		数组、切片、字符串返回索引和值。
		map返回键和值。
		通道（channel）只返回通道内的值。

5. switch case
	Go语言规定每个switch只能有一个default分支。一个分支可以有多个值，多个case值中间使用英文逗号分隔。
	分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量
	fallthrough语法可以执行满足条件的case的下一个case
*/

func ifDemo() {
	// 作用域
	// age变量此时只在if条件判断语句中生效
	if age := 19; age > 18 { // 如果 age > 18 就执行这个{}中的代码
		fmt.Println("澳门首家线上赌场开业啦！")
	} else { // 否则就执行这个{}中的代码
		fmt.Println("改写暑假作业啦！")
	}

	// fmt.Println(age) // 在这里是找不到age
}

func forDemo1() {
	// 基本格式
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 变种1
	// var i = 5
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 变种2
	// var i = 5
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 无限循环
	// for {
	// 	fmt.Println("123")
	// }

	// for range循环
	s := "Hello沙河"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}

func forDemo2() {
	// 1. goto: goto + label 实现跳出多层for循环
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签。跳到我指定的那个标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return

breakTag: // label标签
	fmt.Println("结束for循环")

	// 2. break
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")

	// 3. continue
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}

func switchDemo() {
	// 一个分支可以有多个值，多个case值中间使用英文逗号分隔。
	// 变量可定义在 switch 里
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}

	// 分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。例如：
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

	// fallthrough语法可以执行满足条件的case的下一个case
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}

}
