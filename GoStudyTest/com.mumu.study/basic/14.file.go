// @Author liuzhen
// @Date 2023/12/29 18:13:00
// @Desc
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/**
文件操作

1. 打开和关闭文件
	1.1.os.Open() 打开一个文件，返回一个*File和一个err。对得到的文件实例调用 close() 方法能够关闭文件。
	1.2. defer file.Close()  关闭文件
2. 读取文件: file.Read(b []byte) (n int, err error)  接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回0和io.EOF。

3. bufio 读取文件: bufio是在file的基础上封装了一层API，支持更多的功能。
	3.1. reader := bufio.NewReader(file)  获取 bufio.Reader
	3.2. reader.ReadString('\n')  读取数据

4. 读取整个文件
io/ioutil包的ReadFile方法能够读取完整的文件，只需要将文件名作为参数传入。

5. 文件写入操作
	os.OpenFile() 函数能够以指定模式打开文件，从而实现文件写入相关功能。
		func OpenFile(name string, flag int, perm FileMode) (*File, error) {
			...
		}
		其中:
		name：要打开的文件名 flag：打开文件的模式。 模式有以下几种：s
		模式	含义
		os.O_WRONLY	只写
		os.O_CREATE	创建文件
		os.O_RDONLY	只读
		os.O_RDWR	读写
		os.O_TRUNC	清空
		os.O_APPEND	追加
		perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。

*/

func main() {
	// 1.1. 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}

	// 1.2. 关闭文件
	//file.Close()
	// 为了防止文件忘记关闭，我们通常使用defer注册文件关闭语句。
	defer file.Close()

	// 2. 使用Read方法读取数据
	// 2.1. 读一次
	onceRead(file)
	// 2.2. 循环读取
	forRead(file)
}

// 2.1. 读一次
func onceRead(file *os.File) {
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))
}

// 2.2. 循环读取 使用for循环读取文件中的所有数据。
func forRead(file *os.File) {
	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))

}

// 3. bufio 读取文件
func bufioRead(file *os.File) {
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}

}

// 4. 读取整个文件
func readAll() {
	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))

}
