// @Author liuzhen
// @Date 2023/12/19 20:07:00
// @Desc
package main

import "fmt"

/**
map 容器
map是一种无序的基于key-value的数据结构，Go语言中的map是“引用类型”，必须初始化才能使用。

1. map
	1.1 定义未初始化: map[KeyType]ValueType
				其中: KeyType: 表示键的类型。
				  	  ValueType: 表示键对应的值的类型。

	1.2. 定义且初始化: map类型的变量默认初始值为nil，需要使用make()函数来分配内存。
		语法为: make(map[KeyType]ValueType, [cap])
			其中: cap: 表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。

2. 常用操作:
	2.1. 增加元素: map[key] = value		eg: map["age"] = 30
	2.2. 删除键值对: 使用delete()内建函数从map中删除一组键值对，
		delete()函数的格式如下: delete(map, key)		eg: delete(map, "age") // 将age: 从map中删除
							其中: map: 表示要删除键值对的map
								  key:表示要删除的键值对的键
	2.2. 判断某个键是否存在: value, ok := map[key]
	2.3. 遍历: 使用 for range 遍历map
			scoreMap := make(map[string]int)
			for k, v := range scoreMap {
				fmt.Println(k, v)
			}
	2.4. 只想遍历key:
			for k := range scoreMap {
				fmt.Println(k)
			}
		注意： 遍历map时的元素顺序与添加键值对的顺序无关。

3. 元素为map类型的切片
		var mapSlice = make([]map[string]string, 3)

4. 值为切片类型的map
		var sliceMap = make(map[string][]string, 3)
*/

func main() {
	// 1.1 定义未初始化
	var m1 map[string]int
	fmt.Println(m1 == nil) // 还没有初始化（没有在内存中开辟空间）
	// 1.2. 定义且初始化
	m1 = make(map[string]int, 10) // 要估算好该map容量，避免在程序运行期间再动态扩容

	// 2.1. 增加元素
	m1["理想"] = 18
	m1["jiwuming"] = 35

	fmt.Println(m1)
	fmt.Println(m1["理想"])
	// 约定成俗用ok接收返回的布尔值
	fmt.Println(m1["娜扎"]) // 如果不存在这个key拿到对应值类型的零值

	// 2.2. 判断某个键是否存在
	value, ok := m1["娜扎"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}

	// 2.3. map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 只遍历key
	for k := range m1 {
		fmt.Println(k)
	}
	// 只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}

	// 2.2. 删除键值对
	delete(m1, "jiwuming")
	fmt.Println(m1)
	delete(m1, "沙河") // 删除不存在的key

	//sliceMapDemo()
}

func sliceMapDemo() {
	// 1. 定义元素类型为map的切片
	var s1 = make([]map[int]string, 10, 10)
	// 没有对内部的map做初始化
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "沙河"
	fmt.Println(s1)

	// 2. 定义值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30}
	fmt.Println(m1)
}
