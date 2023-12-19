// @Author liuzhen
// @Date 2023/12/19 22:55:00
// @Desc
package main

import "fmt"

/**
接口(interface): 定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。

1. 接口类型: 接口是一种由程序员来定义的类型，一个接口类型就是一组方法的集合，它规定了需要实现的所有方法。相较于使用结构体类型，当我们使用接口类型说明相比于它是什么更关心它能做什么。
	1.1. 接口的定义: 每个接口类型由任意个方法签名组成，接口的定义格式如下：
		type 接口类型名 interface {
			方法名1(参数列表1) 返回值列表1
			方法名2(参数列表2) 返回值列表2
			...
		}
		其中:
			接口类型名：Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有关闭操作的接口叫closer等。接口名最好要能突出该接口的类型含义。
			方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
			参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

2. 实现接口的条件: 接口就是规定了一个需要实现的方法列表，在 Go 语言中一个类型只要实现了接口中规定的所有方法，那么我们就称它实现了这个接口。

3. 面向接口编程

4. 值接收者和指针接收者
	4.1. 值接收者实现接口: 使用值接收者实现接口之后，不管是结构体类型还是对应的结构体指针类型的变量都可以赋值给该接口变量。
	4.2. 指针接收者实现接口: 只有对应的结构体指针类型的变量才可以赋值给该接口变量。
	由于Go语言中有对指针求值的语法糖，对于值接收者实现的接口，无论使用值类型还是指针类型都没有问题。但是我们并不总是能对一个值求址，所以对于指针接收者实现的接口要额外注意。

5. 类型与接口的关系
	5.1. 一个类型实现多个接口: 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。例如狗不仅可以叫，还可以动。我们完全可以分别定义
	5.2. 多种类型实现同一接口: Go语言中不同的类型还可以实现同一接口。例如在我们的代码世界中不仅狗可以动，汽车也可以动。我们可以使用如下代码体现这个关系。

6. 接口组合
	接口与接口之间可以通过互相嵌套形成新的接口类型，例如Go标准库io源码中就有很多接口之间互相组合的示例。

7. 空接口
	7.1. 空接口的定义: 空接口是指没有定义任何方法的接口类型。因此任何类型都可以视为实现了空接口。也正是因为空接口类型的这个特性，空接口类型的变量可以存储任意类型的值。
		eg: type Any interface{}
	7.2. 通常我们在使用空接口类型时不必使用type关键字声明，可以像下面的代码一样直接使用interface{}。eg: var x interface{}  // 声明一个空接口类型变量x
	7.3. 空接口的应用
		7.3.1. 空接口作为函数的参数: 使用空接口实现可以接收任意类型的函数参数。
		7.3.2. 空接口作为map的值: 使用空接口实现可以保存任意值的字典。

7. 接口值: TODO

8. 类型断言: 接口值可能赋值为任意类型的值，那我们如何从接口值获取其存储的具体数据呢？
	想要从接口值中获取到对应的实际值需要使用类型断言，
	其语法格式如下: v, ok := x.(T)
			其中:
				x: 表示接口类型的变量
				T: 表示断言x可能是的类型。
			该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。

*/

// 7. 空接口
// Any 不包含任何方法的空接口类型
type Any interface{}

// Sayer 定义一个接口类型
type Sayer interface {
	Say()
}

// Mover 定义一个接口类型
type Mover interface {
	Move()
}

func (d Dog) Say() {
	fmt.Println("汪汪汪~")
}

// Move 使用值接收者定义Move方法实现Mover接口
func (d Dog) Move() {
	fmt.Println("狗会动")
}

type Cat struct{}

func (c Cat) Say() {
	fmt.Println("喵喵喵~")
}

// Move 使用指针接收者定义Move方法实现Mover接口
func (c *Cat) Move() {
	fmt.Println("猫会动")
}

// 定义一个通用的 MakeHungry函数，接收 Sayer 类型的参数。
// MakeHungry 饿肚子了...
func MakeHungry(s Sayer) {
	s.Say()
}

func main() {
	var c Cat
	MakeHungry(c)
	var d Dog
	MakeHungry(d)
}

// 4.1. 值类型接收者 demo
func valueReceiverDemo() {
	var mover Mover // 声明一个Mover类型的变量x

	var d1 = Dog{} // d1是Dog类型
	mover = d1     // 可以将d1赋值给变量x
	mover.Move()

	var d2 = &Dog{} // d2是Dog指针类型
	mover = d2      // 也可以将d2赋值给变量x
	mover.Move()

}

// 4.2. 指针类型接收者 demo
func pointerReceiverDemo() {
	var mover Mover // 声明一个Mover类型的变量x

	// 此时实现Mover接口的是*Cat类型，我们可以将*Cat类型的变量直接赋值给Mover接口类型的变量x。
	var c1 = &Cat{} // c1是*Cat类型
	mover = c1      // 可以将c1当成Mover类型
	mover.Move()

	// 但是不能给将Cat类型的变量赋值给Mover接口类型的变量x。
	// 下面的代码无法通过编译！！！
	// var c2 = Cat{} // c2是Cat类型
	// x = c2         // 不能将c2当成Mover类型
}

// 7.3.1. 空接口作为函数参数
func emptyInterfaceDemo1(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

// 7.3.2. 空接口作为map的值
func emptyInterfaceDemo2() {
	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)

}

// 8.1. 断言
func assertDemo1() {
	var n Mover = &Dog{Name: "旺财"}
	v, ok := n.(*Dog)
	if ok {
		fmt.Println("类型断言成功")
		v.Name = "富贵" // 变量v是*Dog类型
	} else {
		fmt.Println("类型断言失败")
	}
}

// 8.2. 对传入的空接口类型变量x进行类型断言
func assertDemo2(x interface{}) {
	// 如果对一个接口值有多个实际类型需要判断，推荐使用switch语句来实现。
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}

}
