package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var filePath = "D:\\data\\GoWorkSpace\\src\\github.com.qq4215279\\StudyGo\\TeaCode\\day05\\12file_open\\read.txt"

// 打开文件
func readFromFile() {
	fileObj, err := os.Open("D:\\data\\GoWorkSpace\\src\\github.com.qq4215279\\StudyGo\\TeaCode\\day05\\12file_open\\read.txt")
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 读文件
	// var tmp = make([]byte, 128) // 指定读的长度
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

// 利用bufIo这个包读取文件
func readFromFileByBufIo() {
	fileObj, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed, err:%v", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromFileByIoUtil() {

	ret, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}
func main() {
	//readFromFile()
	//readFromFileByBufIo()
	readFromFileByIoUtil()

}
