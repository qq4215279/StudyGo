// @Author liuzhen
// @Date 2023/12/20 16:38:00
// @Desc
package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

/**
测试(Test)

1. go test工具
Go语言中的测试依赖 go test 命令。编写测试代码和编写普通的Go代码过程是类似的，并不需要学习新的语法、规则或工具。
go test 命令是一个按照一定约定和组织的测试代码的驱动程序。在包目录内，所有以 _test.go 为后缀名的源代码文件都是 go test 测试的一部分，不会被 go build 编译到最终的可执行文件中。
go test 命令会遍历所有的 *_test.go 文件中符合上述命名规则的函数，然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。
go test 命令:
	go test 命令添加 -v 参数，查看测试函数名称和运行时间。eg: go test -v
	go test 命令后添加 -run 参数，它对应一个正则表达式，只有函数名匹配上的测试函数才会被go test命令执行。eg: go test -v -run="More"
		还可以通过 / 来指定要运行的子测试用例，例如TestSubSplit(): go test -v -run=SubSplit/simple 只会运行simple对应的子测试用例。
	go test 命令后添加 -cover 参数，来查看测试覆盖率。eg: go test -cover
		-coverprofile 参数，用来将覆盖率相关的记录信息输出到一个文件。eg: go test -cover -coverprofile=c.out
		在c.out 目录中执行: go tool cover -html=c.out，使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告。报告中绿色标记的语句块表示被覆盖了，而红色的表示没有被覆盖。


在*_test.go文件中有三种类型的函数，单元测试函数、基准测试函数和示例函数。
	类型			格式							作用
	测试函数	函数名前缀为Test				测试程序的一些逻辑行为是否正确
	基准函数	函数名前缀为Benchmark		测试函数的性能
	示例函数	函数名前缀为Example			为文档提供示例文档

2. 测试函数
	2.1. 每个测试函数必须导入testing包，测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头。测试函数的基本格式（签名）如下:
			func TestName(t *testing.T){
				// ...
			}
		其中参数t用于报告测试失败和附加的日志信息。 testing.T 的拥有的方法如下:
			func (c *T) Error(args ...interface{})
			func (c *T) Errorf(format string, args ...interface{})
			func (c *T) Fail()
			func (c *T) FailNow()
			func (c *T) Failed() bool
			func (c *T) Fatal(args ...interface{})
			func (c *T) Fatalf(format string, args ...interface{})
			func (c *T) Log(args ...interface{})
			func (c *T) Logf(format string, args ...interface{})
			func (c *T) Name() string
			func (t *T) Parallel()
			func (t *T) Run(name string, f func(t *T)) bool
			func (c *T) Skip(args ...interface{})
			func (c *T) SkipNow()
			func (c *T) Skipf(format string, args ...interface{})
			func (c *T) Skipped() bool

	2.2. 测试组: 用于添加更多的测试用例
	2.3. 子测试: 如果测试用例比较多的时候，我们是没办法一眼看出来具体是哪个测试用例失败了，则使用子测试
	2.4. 测试覆盖率

3. 基准测试
3.1. 基准测试就是在一定的工作负载之下检测程序性能的一种方法。基准测试的基本格式如下:
		func BenchmarkName(b *testing.B){
			// ...
		}
	基准测试并不会默认执行，需要增加 -bench 参数，所以我们通过执行 go test -bench=Split 命令执行基准测试。
		为基准测试添加-benchmem参数，来获得内存分配的统计数据。eg: go test -bench=Split -benchmem
	基准测试以Benchmark为前缀，需要一个*testing.B类型的参数b，基准测试必须要执行b.N次，这样的测试才有对照性，b.N的值是系统根据实际情况去调整的，从而保证测试的稳定性。
	testing.B拥有的方法如下:
		func (c *B) Error(args ...interface{})
		func (c *B) Errorf(format string, args ...interface{})
		func (c *B) Fail()
		func (c *B) FailNow()
		func (c *B) Failed() bool
		func (c *B) Fatal(args ...interface{})
		func (c *B) Fatalf(format string, args ...interface{})
		func (c *B) Log(args ...interface{})
		func (c *B) Logf(format string, args ...interface{})
		func (c *B) Name() string
		func (b *B) ReportAllocs()
		func (b *B) ResetTimer()
		func (b *B) Run(name string, f func(b *B)) bool
		func (b *B) RunParallel(body func(*PB))
		func (b *B) SetBytes(n int64)
		func (b *B) SetParallelism(p int)
		func (c *B) Skip(args ...interface{})
		func (c *B) SkipNow()
		func (c *B) Skipf(format string, args ...interface{})
		func (c *B) Skipped() bool
		func (b *B) StartTimer()
		func (b *B) StopTimer()

	3.2. 性能能比较函数
	上面的基准测试只能得到给定操作的绝对耗时，但是在很多性能问题是发生在两个不同操作之间的相对耗时，比如同一个函数处理1000个元素的耗时与处理1万甚至100万个元素的耗时的差别是多少？再或者对于同一个任务究竟使用哪种算法性能最佳？我们通常需要对两个不同算法的实现使用相同的输入来进行基准比较测试。
	性能比较函数通常是一个带有参数的函数，被多个不同的Benchmark函数传入不同的值来调用。

	3.3. 重置时间
		b.ResetTimer之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作。

	3.4. 并行测试
		func (b *B) RunParallel(body func(*PB))会以并行的方式执行给定的基准测试。
		RunParallel会创建出多个goroutine，并将b.N分配给这些goroutine执行， 其中goroutine数量的默认值为GOMAXPROCS。
		用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性， 那么可以在RunParallel之前调用SetParallelism 。RunParallel通常会与-cpu标志一同使用。

4. Setup 与 TearDown: 测试程序有时需要在测试之前进行额外的设置(setup)或在测试之后进行拆卸(teardown)。
	4.1. TestMain: 通过在 *_test.go 文件中定义 TestMain 函数来可以在测试之前进行额外的设置(setup)或在测试之后进行拆卸（teardown）操作。
		如果测试文件包含函数: func TestMain(m *testing.M) 那么生成的测试会先调用 TestMain(m)，然后再运行具体测试。
		TestMain运行在主goroutine中, 可以在调用m.Run前后做任何设置(setup)和拆卸(teardown)。退出测试的时候应该使用m.Run的返回值作为参数调用 os.Exit。

	4.2. 子测试的Setup与Teardown
	有时候我们可能需要为每个测试集设置Setup与Teardown，也有可能需要为每个子测试设置Setup与Teardown。

5. 示例函数
	被go test特殊对待的第三种函数就是示例函数，它们的函数名以Example为前缀。它们既没有参数也没有返回值。标准格式如下：
		func ExampleName() {
			// ...
		}
	示例函数用处:
		示例函数能够作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联。
		示例函数只要包含了// Output:也是可以通过go test运行的可执行测试。
		示例函数提供了可以直接运行的示例代码，可以直接在golang.org的godoc文档服务器上使用Go Playground运行示例代码。
*/

// 2.1.1 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
// 在当前包路径下，执行 go test 命令，可以看到输出结果
func TestSplit(t *testing.T) { //
	// 程序输出的结果
	got := Split("a:b:c", ":")
	// 期望的结果
	want := []string{"a", "b", "c"}
	// 因为slice不能比较直接，借助反射包中的方法比较
	if !reflect.DeepEqual(want, got) {
		// 测试失败输出错误提示
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// 2.1.2 一个测试用例有点单薄，我们再编写一个测试使用多个字符切割字符串的例子。再次运行go test命令
func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// 2.1 Split() 测试
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		// 错误写法
		//s = s[i+1:]
		// 正确写法: 这里使用len(sep)获取sep的长度
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

// 2.2. 测试组。把多个测试用例合到一起，再次执行 go test 命令。
func TestGroupSplit(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	// 定义一个存储测试用例的切片
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	// 遍历切片，逐一执行测试用例
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			// t.Errorf("expected:%v, got:%v", tc.want, got)
			// 打印的测试失败提示信息：expected:[河有 又有河], got:[ 河有 又有河]，你会发现[ 河有 又有河]中有个不明显的空串，
			// 这种情况下十分推荐使用 %#v 的格式化方式。如下:
			t.Errorf("expected:%#v, got:%#v", tc.want, got)
		}
	}
}

// 2.3. 子测试
func TestSubSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}

	// 方式1:
	/*for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}*/

	// 4.2.1. 测试之前执行setup操作
	teardownTestCase := setupTestCase(t)
	// 4.2.1. 测试之后执行testdoen操作
	defer teardownTestCase(t)

	// 方式2: Go1.7+中新增了子测试，我们可以按照如下方式使用t.Run执行子测试
	for name, tc := range tests {
		// 使用t.Run()执行子测试
		t.Run(name, func(t *testing.T) {
			// 4.2.2 子测试之前执行setup操作
			teardownSubTest := setupSubTest(t)
			// 4.2.2. 测试之后执行testdoen操作
			defer teardownSubTest(t)
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

// 3.1. 基准测试 示例
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

// 3.2. 性能比较函数
func benchmark(b *testing.B, size int) {
	/* ... */
}

// 3.2. 运行基准测试： go test -bench=.
func Benchmark10(b *testing.B)   { benchmark(b, 10) }
func Benchmark100(b *testing.B)  { benchmark(b, 100) }
func Benchmark1000(b *testing.B) { benchmark(b, 1000) }

// 3.3. 重置时间
func BenchmarkResetTimeSplit(b *testing.B) {
	time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
	b.ResetTimer()              // 重置计时器
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

// 3.4. 并行测试。运行基准测试: go test -bench=.
// 还可以通过在测试命令后添加-cpu参数如go test -bench=. -cpu 1来指定使用的CPU数量。
func BenchmarkSplitParallel(b *testing.B) {
	// b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("沙河有沙又有河", "沙")
		}
	})
}

// 4.1. 一个使用TestMain来设置Setup和TearDown的示例如下：
func TestMain(m *testing.M) {
	// 测试之前的做一些设置
	fmt.Println("write setup code here...")
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()

	// 执行测试
	retCode := m.Run()

	// 测试之后做一些拆卸工作
	fmt.Println("write teardown code here...")
	// 退出测试
	os.Exit(retCode)
}

// 4.2.1. 测试集的 Setup 与 Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:测试之后的teardown")
	}
}

// 4.2.2. 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:子测试之后的teardown")
	}
}

// 5. 示例函数示例
func ExampleSplit() {
	fmt.Println(Split("a:b:c", ":"))
	fmt.Println(Split("沙河有沙又有河", "沙"))
	// Output:
	// [a b c]
	// [ 河有 又有河]
}
