// @Author liuzhen
// @Date 2023/12/19 19:39:00
// @Desc
package main

import "fmt"

/**
指针
& 取地址
* 根据地址取值

1. 指针地址和指针类型
	每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用&字符放在变量前面对变量进行“取地址”操作。
	Go语言中的值类型: int、float、bool、string、array、struct 都有对应的指针类型，如：*int、*int64、*string等。
	Go语言中的引用类型: slice、map、channel
	取变量指针的语法如下: ptr := &v    // v的类型为T
		其中: v: 代表被取地址的变量，类型为T
		  	  ptr: 用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针。

2. 指针传值/取值: 在对普通变量使用&操作符取地址后会获得这个变量的指针，然后可以对指针使用*操作，也就是指针传值/取值
	eg: a := 10
		b := &a // 取变量a的地址，将指针保存到b中
		*b = 100 // 指针传值
		value := *b // 指针取值

3. new 和 make
	在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
	而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
	要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存。
	3.1. new
		签名如下: func new(Type) *Type
			eg: a := new(int)
		其中: Type表示类型，new函数只接受一个参数，这个参数是一个类型
			  *Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
		new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。

	3.2. make
		签名如下: func make(t Type, size ...IntegerType) Type
			eg: var b = make(map[string]int, 10)
		make也是用于内存分配的，区别于new，它只用于slice、map以及channel的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，
		因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
		make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。

4. new与make的区别
	二者都是用来做内存分配的。
	make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。

*/

func main() {
	// 1. &:取地址
	// n := 18
	// p := &n
	// fmt.Println(p)
	// fmt.Printf("%T\n", p) // *int：int类型的指针

	// 2. *：根据地址取值
	// m := *p
	// fmt.Println(m)
	// fmt.Printf("%T\n", m) // int

	var a1 *int // nil pointer
	fmt.Println(a1)
	// new函数申请一个内存地址
	var a2 = new(int)
	fmt.Println(a2)
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(*a2)

	// 使用make初始化map
	var b = make(map[string]int, 10)
	fmt.Println(b)
}

func panicDemo() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["沙河娜扎"] = 100
	fmt.Println(b)

}
