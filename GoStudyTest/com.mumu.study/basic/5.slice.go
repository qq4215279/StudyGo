// @Author liuzhen
// @Date 2023/12/19 18:58:00
// @Desc
package main

import (
	"fmt"
	"sort"
)

/**
切片
切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
切片是一个“引用类型”，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
	1. 切片不保存具体的值
	2. 切片对应一个底层数组
	3. 底层数组都是占用一块连续的内存

1. 切片的定义
	声明切片类型的基本语法如下: var name []T
		name: 表示变量名
		T: 表示切片中的元素类型
	eg: 声明一个字符串切片: var a []string
		声明一个整型切片并初始化: var b = []int{}

2. 切片的长度和容量
	切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。

3. 切片表达式
	切片表达式从字符串、数组、指向数组或切片的指针构造子字符串或切片。它有两种变体：一种指定low和high两个索引界限值的简单的形式，另一种是除了low和high索引界限值外还指定容量的完整的形式。

	3.1. 简单切片表达式  a = a[low : high]
		切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。 切片表达式中的low和high表示一个索引范围（左包含，右不包含），
		也就是下面代码中从数组a中选出 1<=索引值<4 的元素组成切片s，得到的切片长度 = high - low，容量等于得到的切片的底层数组的容量。
		注: 对于数组或字符串，如果 0 <= low <= high <= len(a)，则索引合法，否则就会索引越界（out of range）。
		eg: a[2:]  // 等同于 a[2:len(a)]
			a[:3]  // 等同于 a[0:3]
			a[:]   // 等同于 a[0:len(a)]

	3.2. 完整切片表达式。对于数组，指向数组的指针，或切片a(注意不能是字符串)支持完整切片表达式: a = a[low : high : max]
		完整切片表达式需要满足的条件是 0 <= low <= high <= max <= cap(a)，其他条件和简单切片表达式相同。

4. 使用 make() 函数构造切片
	我们上面都是基于数组来创建的切片，如果需要动态的创建一个切片，我们就需要使用内置的 make() 函数，
	格式如下: make([]T, size, cap)
		其中: T: 切片的元素类型
			  size: 切片中元素的数量
			  cap: 切片的容量
	eg: a := make([]int, 2, 10)

5. 切片的本质
	切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。

6. 常用操作：
	6.1. 切片判空: 判断切片是否为空，要检查切片是否为空，请始终使用 len(s) == 0 来判断，而不应该使用 s == nil 来判断。
	6.2. 切片比较: 切片之间是不能比较的，我们不能使用 == 操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和nil比较。
		一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是nil
	6.3. 切片的赋值拷贝: 拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容
		eg: s1 := make([]int, 3) // [0 0 0]
			s2 := s1             // 将s1直接赋值给s2，s1和s2共用一个底层数组
			s2[0] = 100 		 // s1 = s2 = [100 0 0]
	6.4. 切片遍历: 切片的遍历方式和数组是一致的，支持索引遍历和 for range 遍历。
		eg: s := []int{1, 3, 5}
			// 方式1:
			for i := 0; i < len(s); i++ {
				fmt.Println(i, s[i])
			}
			// 方式2:
			for index, value := range s {
				fmt.Println(index, value)
			}

	6.5. 添加元素: append() 调用append函数必须用原来的切片变量接收返回值，append追加元素，原来的底层数组放不下的时候，Go底层就会把底层数组换一个
		eg: var s []int
			s = append(s, 1)        // [1]
			s = append(s, 2, 3, 4)  // [1 2 3 4]
			s2 := []int{5, 6, 7}
			s = append(s, s2...)    // [1 2 3 4 5 6 7]
		每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，
		此时该切片指向的底层数组就会更换。“扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值。
		切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。

	6.6. 删除元素: append()。即: 要从切片a中删除索引为index的元素，操作方法是 a = append(a[:index], a[index+1:]...)
		eg: a := []int{30, 31, 32, 33, 34, 35, 36, 37}
			a = append(a[:2], a[3:]...) // 要删除索引为2的元素

	6.7. 复制切片: copy(destSlice, srcSlice []T)
				其中: srcSlice: 数据来源切片
				  	  destSlice: 目标切片

	6.8 切片排序: sort包: sort.Ints(s)
		eg: s := []int{3, 7, 8, 9, 1}
			sort.Ints(s)


切片的扩容策略:
首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，
	即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。
*/

func main() {
	// 1.1 切片的定义
	var s1 []int    // 定义一个存放int类型元素的切片
	var s2 []string // 定义一个存放string类型元素的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // true
	fmt.Println(s2 == nil) // true

	// 1.2. 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张江", "平山村"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // false
	fmt.Println(s2 == nil) // false

	// 2. 长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 3. 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] // 基于一个数组切割，左包含右不包含，（左闭右开）
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)
	s5 := a1[:4] // => [0:4] [1 3 5 7]
	s6 := a1[3:] // => [3:len(a1)]  [7 9 11 13]
	s7 := a1[:]  // => [0:len(a1)]
	fmt.Println(s5, s6, s7)

	// 4. 切片的容量是指底层数组的容量
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
	// 底层数组从切片的第一个元素到最后的元素数量
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6))

	// 5. 切片再切割
	s8 := s6[3:] // [13]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8))

	// 6. 切片是引用类型，都指向了底层的一个数组。
	fmt.Println("s6：", s6)
	a1[6] = 1300 // 修改了底层数组的值
	fmt.Println("s6：", s6)
	fmt.Println("s8：", s8)

	// 7. make()函数创造切片
	s11 := make([]int, 5, 10)
	fmt.Printf("s11=%v len(s11)=%d cap(s11)=%d\n", s11, len(s11), cap(s11))
	s22 := make([]int, 0, 10)
	fmt.Printf("s22=%v len(s22)=%d cap(s22)=%d\n", s22, len(s22), cap(s22))

	// 8. 切片的赋值
	s33 := []int{1, 3, 5}
	s44 := s33 // s33 和 s44 都指向了同一个底层数组
	fmt.Println(s33, s44)
	s33[0] = 1000
	fmt.Println(s33, s44)

	// 9. 切片排序
	a99 := []int{3, 7, 8, 9, 1}
	fmt.Println("a99排序前: ", a99)
	sort.Ints(a99) // 对切片进行排序
	fmt.Println("a99排序后: ", a99)

	// 切片的遍历
	// rangeDemo()

	// 切片拷贝
	// copyDemo()
}

func appendDemo() {
	s1 := []string{"北京", "上海", "深圳"}
	// 添加一个元素
	s1 = append(s1, "广州")
	// 添加多个元素
	s1 = append(s1, "杭州", "成都")

	// 添加一个切片
	ss := []string{"武汉", "西安", "苏州"}
	s1 = append(s1, ss...) // ...表示拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	// 删掉索引为1的那个“"上海"”
	s1 = append(s1[0:1], s1[2:]...)
	fmt.Println(s1)
}

// 切片的遍历
func rangeDemo() {
	s3 := []int{1, 3, 5}
	// 1. 索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	// 2. for range循环
	for i, v := range s3 {
		fmt.Println(i, v)
	}

}

// 切片拷贝
func copyDemo() {
	// 1. 切片不保存具体的值
	// 2. 切片对应一个底层数组
	// 3. 底层数组都是占用一块连续的内存
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(b) //[1 2 3 4 5]
	// 修改底层数组
	b[0] = 1000
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(b) //[1000 2 3 4 5]

	fmt.Println("------------------->")

	// copy()复制切片
	aa := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	copy(c, aa)             // 使用copy()函数将切片a中的元素复制到切片c
	fmt.Println("aa: ", aa) // [1 2 3 4 5]
	c[0] = 1000
	fmt.Println("aa: ", aa) //[1 2 3 4 5]
	fmt.Println("c: ", c)   //[1000 2 3 4 5]
}
