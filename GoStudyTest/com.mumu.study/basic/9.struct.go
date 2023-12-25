// @Author liuzhen
// @Date 2023/12/19 21:47:00
// @Desc
package main

import (
	"encoding/json"
	"fmt"
)

/**
结构体
Go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，这时候再用单一的基本数据类型明显就无法满足需求了。
Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体，英文名称struct。

1. 自定义类型 和 类型别名
	1.1. 自定义类型: 自定义类型是定义了一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义。
		eg: type MyInt int // 将MyInt定义为int类型。通过type关键字的定义，MyInt就是一种新的类型，它具有int的特性。
	1.2. 类型别名: 类型别名规定: TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。就像一个孩子小时候有小名、乳名，上学后用学名，英语老师又会给他起英文名，但这些名字都指的是他本人。
		eg: type TypeAlias = Type
			type byte = uint8
			type rune = int32

2. 结构体的定义: 使用type和struct关键字来定义结构体，具体代码格式如下：
			type 类型名 struct {
				字段名 字段类型
				字段名 字段类型
				…
			}
		其中:
			类型名: 标识自定义结构体的名称，在同一个包内不能重复。
			字段名: 表示结构体字段名。结构体中的字段名必须唯一。
			字段类型: 表示结构体字段的具体类型。
		eg: type person struct {
				name, city string
				age        int8
			}

3. 结构体实例化: 只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
	3.1. 基本实例化:
			结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。格式: var 结构体实例 结构体类型
			通过 . 来访问结构体的字段（成员变量）
			eg: var p1 person	p1.name = "沙河娜扎"  p1.city = "北京"
	3.2. 匿名结构体: eg: var user struct{Name string; Age int}
    			   		 user.Name = "小王子"
	3.3. 创建指针类型结构体: 我们还可以通过使用 new关键字对结构体进行实例化，得到的是结构体的地址。
			eg: var p2 = new(person) // *main.person
	3.4. 取结构体的地址实例化: 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
			eg: p3 := &person{}  // *main.person
			注: p3.name = "七米"  其实在底层是 (*p3).name = "七米"，这是Go语言帮我们实现的语法糖。

4. 结构体初始化: 没有初始化的结构体，其成员变量都是对应其类型的零值。eg: var p4 person  => //p4=main.person{name:"", city:"", age:0}
	4.1. 使用键值对初始化: p5 := person{ name: "小王子", city: "北京", age:  18,}
	4.2. 对结构体指针进行键值对初始化: p6 := &person{city: "北京", age:  18,} // p6=&main.person{name:"", city:"北京", age:18}
		注: 当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
	4.3. 使用值的列表初始化: p7 := &person{"沙河娜扎", "北京", 28,} // p7=&main.person{name:"沙河娜扎", city:"北京", age:28}
		使用这种格式初始化时，需要注意：
			必须初始化结构体的所有字段。
			初始值的填充顺序必须与字段在结构体中的声明顺序一致。
			该方式不能和键值初始化方式混用。

5. 结构体内存布局: 结构体占用一块连续的内存。

6. 空结构体: 空结构体是不占用空间的。
	eg: var v struct{}   fmt.Println(unsafe.Sizeof(v)) // 0

7. 构造函数: Go语言的结构体没有构造函数，我们可以自己实现。eg: func newPerson(name, city string, age int8) *person {...}
	- 构造函数: 约定成俗用new开头
	- 返回的是结构体还是结构体指针
	- 当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销

8. 方法和接收者: 方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。
	8.1. 方法的定义格式如下:
		func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
			函数体
		}
		其中:
			接收者变量: 接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是self、this之类的命名。
				例如: Person类型的接收者变量应该命名为p，Connector类型的接收者变量应该命名为c等。
			接收者类型: 接收者类型和参数类似，可以是指针类型和非指针类型。
			方法名、参数列表、返回参数：具体格式与函数定义相同。
	方法与函数的区别是，函数不属于任何类型，方法属于特定的类型！！
	8.2. 值类型的接收者(传拷贝进去): 当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
	8.3. 指针类型的接收者(传内存地址进去): 指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。这种方式就十分接近于其他语言中面向对象中的this或者self。
	什么时候应该使用指针类型接收者
		需要修改接收者中的值
		接收者是拷贝代价比较大的大对象
		保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

9. 任意类型添加方法
	在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
	举个例子，我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。
	注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法（不能给别的包里面的类型添加方法,只能给自己包里的类型添加方法）。

10. 结构体的匿名字段: 结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。eg: noNameOb
	注意：这里匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。

11. 结构体的“继承”: 通过嵌套匿名结构体实现继承
	嵌套结构体: 一个结构体中可以嵌套包含另一个结构体或结构体指针

12. 结构体字段的可见性: 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

13. 结构体与JSON序列化: 使用 json包中的api。注: 私有不能被json包访问!
	13.1. JSON序列化(结构体 -> JSON格式的字符串): jsonStr, err := json.Marshal(strcut)
	13.2. JSON反序列化(JSON格式的字符串 -> 结构体): strcut := &person{}
												   err = json.Unmarshal([]byte(str), strcut)

14. 结构体标签(Tag): Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。 Tag在结构体字段的后方定义，由一对反引号包裹起来，
	具体的格式如下： `key1:"value1" key2:"value2"`
	结构体tag由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。同一个结构体字段可以设置多个键值对tag，不同的键值对之间使用空格分隔。
	注意事项: 为结构体编写Tag时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，
		一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。例如不要在key和value之间添加空格。

*/
// 1.1. 自定义类型
type NewInt int

// 1.2. 类型别名
type MyInt = int

// 定义一个结构体
type Person struct {
	Name   string   `json:"id" db:"name" ini:"name"` // 通过指定tag实现json序列化该字段时的key
	City   string   // json序列化是默认使用字段名作为key
	age    int8     // 私有不能被json包访问
	dreams []string // 私有不能被json包访问
}

// 10. 结构体的匿名字段
type noNameOb struct {
	string
	int
}

// 11. 嵌套结构体 address 地址结构体
type address struct {
	Province string
	City     string
}

// 构造函数
func newPerson(Name, City string, age int8) *Person {
	// 问题: 为什么要有构造函数
	// 别人调用我,我能给她一个person类型的变量
	return &Person{
		Name: Name,
		City: City,
		age:  age,
	}
}

// Dream Person 做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.Name)
}

func (p *Person) SetDreams(dreams []string) {
	p.dreams = dreams
}

// 正确的做法是在方法中使用传入的slice的拷贝进行结构体赋值。
func (p *Person) SetDreams2(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

// 8.2. 使用值接收者。go语言中函数传参数永远传的是拷贝
// SetAge2 设置p的年龄
func (p Person) SetAge2(newAge int8) {
	// 修改的是副本的 age
	p.age = newAge
}

// 8.3. 使用指针接收者
// SetAge 设置p的年龄
func (p *Person) SetAge(newAge int8) {
	// 根据内存地址找到那个原变量,修改的就是原来的变量
	// (*p).age = newAge

	// 语法糖,自动根据指针找对应的变量
	p.age = newAge
}

// 9. SayHello 为 NewInt 添加一个 SayHello 的方法
func (m NewInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}

// 12. 结构体的继承
// Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

// Dog 狗
type Dog struct {
	Name    string
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	// 声明一个int32类型的变量x,它的值是10
	// 方法1:
	// var x int32
	// x = 10
	// 方法2:
	// var x int32 = 10
	// 方法3:
	// var x = int32(10)
	// 方法4:
	// x := int32(10)
	// fmt.Println(x)

	// 声明一个myInt类型的变量m,它的值是100
	// 方法1:
	// var m myInt
	// m = 100
	// 方法2:
	// var m myInt = 100
	// 方法3:
	// var m = myInt(100)
	// 方法4
	// m := myInt(100) // 强制类型转换
	// fmt.Println(m)
	// m := myInt(100)
	// m.hello()

	// 3.2. 匿名结构体:多用于临时场景
	var s struct {
		x string
		y int
	}
	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s)

	// 11. 结构体的继承
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ // 注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！

	// 15. 结构体和方法补充知识点
	fmt.Println("结构体和方法补充知识点 ---------------->")
	p1 := Person{Name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？SetDreams这种赋值slice的方式，会导致外部的值改了之后，结构体里的值也发生变化！
	data[1] = "不睡觉"
	fmt.Println(p1.dreams) // ?

	// 13. 序列化 与 反序列化
	serializeDemo()
}

// 序列化 与 反序列化 demo
func serializeDemo() {
	p1 := Person{
		Name: "周林",
		age:  12,
	}

	fmt.Println("serializeDemo start ---------->")

	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("序列化 %v\n", string(b))

	// 反序列化
	str := `{"name":"理想","age":18}`
	var p2 Person
	json.Unmarshal([]byte(str), &p2) // 传指针是为了能在json.Unmarshal内部修改p2的值
	fmt.Printf("反序列化 %#v\n", p2)

	fmt.Println("serializeDemo end ---------->")
}
