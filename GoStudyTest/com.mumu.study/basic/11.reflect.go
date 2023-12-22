// @Author liuzhen
// @Date 2023/12/20 10:31:00
// @Desc
package main

import (
	"fmt"
	"reflect"
)

/**
反射 reflect
在Go语言的反射机制中，任何接口值都由是一个具体类型 和 具体类型的值两部分组成的。
在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为由 reflect.Type 和 reflect.Value 两部分组成，
并且reflect包提供了 reflect.TypeOf 和 reflect.ValueOf 两个函数来获取任意对象的 Type 和 Value。

1. reflect.TypeOf()  函数可以获得任意值的类型对象(reflect.Type)，程序通过类型对象可以访问任意值的类型信息。eg: type := reflect.TypeOf(x)
		在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。因为在Go语言中我们可以使用type关键字构造很多自定义类型。
	1.1. type.Name()  获取类型(Type)。
			注: Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
	1.2. type.Kind()  获取种类(Kind)，就是指底层的类型。但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类(Kind)。
	1.3. 与结构体相关的方法:
		1.3.1. type.NumField() int	返回结构体成员字段数量。
		1.3.2. type.Field(i int) StructField	根据索引，返回索引对应的结构体字段的信息。
		1.3.3. type.FieldByName(name string) (StructField, bool)	根据给定字符串返回字符串对应的结构体字段的信息。
		1.3.4. type.FieldByIndex(index []int) StructField	多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。
		1.3.5. type.FieldByNameFunc(match func(string) bool) (StructField, bool)	根据传入的匹配函数匹配需要的字段。
		1.3.6. type.NumMethod() int	返回该类型的方法集中方法的数目
		1.3.7. type.Method(int) Method	返回该类型方法集中的第i个方法
		1.3.8. type.MethodByName(string)(Method, bool)	根据方法名返回该类型方法集中的方法
	1.4. StructField类型: 用来描述结构体中的一个字段的信息。定义如下:
			type StructField struct {
				// Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
				// 参见http://golang.org/ref/spec#Uniqueness_of_identifiers
				Name    string
				PkgPath string
				Type      Type      // 字段的类型
				Tag       StructTag // 字段的标签
				Offset    uintptr   // 字段在结构体中的字节偏移量
				Index     []int     // 用于Type.FieldByIndex时的索引切片
				Anonymous bool      // 是否匿名字段
			}

2. reflect.ValueOf()  返回的是reflect.Value类型，其中包含了原始值的值信息。reflect.Value与原始值之间可以互相转换。eg: value := reflect.ValueOf(x)
	2.1. 常见值操作:
		2.1.1. value.Interface() interface {}	将值以 interface{} 类型返回，可以通过类型断言转换为指定类型
		2.1.2. value.Int() int64	将值以 int 类型返回，所有有符号整型均可以此方式返回
		2.1.3. value.Uint() uint64	将值以 uint 类型返回，所有无符号整型均可以此方式返回
		2.1.4. value.Float() float64	将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回
		2.1.5. value.Bool() bool	将值以 bool 类型返回
		2.1.6. value.Bytes() []bytes	将值以字节数组 []bytes 类型返回
		2.1.7. value.String() string	将值以字符串类型返回
		2.1.8. value.IsNil() bool  报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。
					常被用于判断指针是否为空；
		2.1.9. value.IsValid() bool  返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。
					常被用于判定返回值是否有效。

3. 通过反射设置变量的值
	想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的 Elem() 方法来获取指针对应的值。
	3.1. value.Elem().Kind()
	3.2. value.Elem().SetInt(val)


反射是一个强大并富有表现力的工具，能让我们写出更灵活的代码。但是反射不应该被滥用，原因有以下三个。
	基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发panic，那很可能是在代码写完的很长时间之后。
	大量使用反射的代码通常难以理解。
	反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个数量级。
*/

type myInt int64

type student struct {
	Name  string `json:"name" zhoulin:"嘿嘿嘿"`
	Score int    `json:"score" zhoulin:"哈哈哈"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

// 1.1 + 1.2 反射获取类型
func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

// 1.3 type 与结构体操作相关的api
func reflectTypeApiDemo() {
	stu := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}

}

// 1.4 StructField类型
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

// 2. 反射获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	// 值的类型种类
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int() 从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float() 从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float() 从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 3. 通过反射设置变量的值
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}

// 3. 通过反射设置变量的值
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200) // 修改的是副本，reflect包会引发panic
	}
}

func main() {
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	reflectType(a) // type: kind:ptr
	reflectType(b) // type:myInt kind:int64
	reflectType(c) // type:int32 kind:int32

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "沙河小王子",
		age:  18,
	}
	var e = book{title: "《跟小王子学Go语言》"}
	reflectType(d) // type:person kind:struct
	reflectType(e) // type:book kind:struct

	// ===========================>

	var aa float32 = 3.14
	var bb int64 = 100
	reflectValue(aa) // type is float32, value is 3.140000
	reflectValue(bb) // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	cc := reflect.ValueOf(10)
	fmt.Printf("type cc :%T\n", cc) // type c :reflect.Value
}
