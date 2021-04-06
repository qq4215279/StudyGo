package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime.Caller()

// 反射api
func f() {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println("funcName: ", funcName) //  main.f1
	fmt.Println("file: ", file)         // D:/data/GoWorkSpace/src/github.com.qq4215279/StudyGo/TeaCode/day06/06runtime_demo/main.go
	fmt.Println(path.Base(file)) // main.go
	fmt.Println("line: ", line) // 25
}

func f1() {
	f()
}

func main() {
	f1()
}
