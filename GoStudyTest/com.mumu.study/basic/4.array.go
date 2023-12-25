// @Author liuzhen
// @Date 2023/12/19 18:10:00
// @Desc
package main

import "fmt"

/**
Array(数组)
	数组是一个“值类型”!
	存放元素的容器
	必须指定存放的元素的类型和容量（长度）
	数组的长度是数组类型的一部分

1. 数组定义: var 数组变量名 [元素数量]T  eg: var a [3]int // 定义一个长度为3元素类型为int的数组a
	注: 数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。 [5]int 和 [10]int 是不同的类型。

2. 数组初始化
	2.1. 使用初始化列表来设置数组元素的值
		数组会初始化为int类型的零值: var testArray [3]int
		使用指定的初始值完成初始化: var numArray = [3]int{1, 2}
	2.2. 让编译器根据初始值的个数自行推断数组的长度。使用: ...
		numArray = [...]int{1, 2}
	2.3. 使用指定索引值的方式来初始化数组
		a := [...]int{1: 1, 3: 5}

3. 数组的遍历
	3.1. for循环遍历
		var a = [...]string{"北京", "上海", "深圳"}
		for i := 0; i < len(a); i++ {
			fmt.Println(a[i])
		}

	3.2. for range遍历
		for index, value := range a {
			fmt.Println(index, value)
		}

4. 多维数组。
	注意： 多维数组只有第一层可以使用...来让编译器推导数组长度。
	eg: a := [...][2]string{
			{"北京", "上海"},
			{"广州", "深圳"},
		}

5. 数组是值类型。数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	注意:
		数组支持 “=="、”!=" 操作符，因为内存总是被初始化过的。
		[n]*T表示指针数组，*[n]T表示数组指针 。

*/

func main() {
	// 1. 数组定义
	var a1 [3]bool // [true false true]
	var a2 [4]bool // [true true false false]
	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	// 2. 数组的初始化
	// 如果不初始化：默认元素都是零值（布尔值：false, 整型和浮点型都是0, 字符串：""）
	fmt.Println(a1, a2)
	// 2.1. 初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)

	// 2.2. 初始化方式2：根据初始值自动推断数组的长度是多少
	// a10 := [9]int{0, 1, 2, 3, 4, 4, 5, 6, 7}
	a10 := [...]int{0, 1, 2, 3, 4, 4, 5, 6, 7}
	fmt.Println(a10)

	// 2.3. 初始化方式3：根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	// 3. 数组的遍历
	citys := [...]string{"北京", "上海", "深圳"} // 索引：0~2 citys[0],citys[1],citys[2]
	// 3.1. 根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	// 3.2. for range遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	// 4. 多维数组
	// [[1 2] [3 4] [5 6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	// 多维数组的遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	fmt.Println("5. ---------------------->")
	// 5. 数组是值类型。数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	b1 := [3]int{1, 2, 3} // [1 2 3]
	modifyArray(b1)       // 在modify中修改的是b1的副本x
	fmt.Println(b1)       // [10 20 30]

	b2 := b1            // [1 2 3] Ctrl+C Ctrl+V => 把world文档从文件夹A拷贝到文件夹B
	b2[0] = 100         // b2:[100 2 3]
	fmt.Println(b1, b2) // b1:[1 2 3]  b2:[100 2 3]
}

func modifyArray(x [3]int) {
	x[0] = 100
}
