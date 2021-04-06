package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// log demo

func main() {
	fileObj, err := os.OpenFile("D:\\data\\GoWorkSpace\\src\\github.com.qq4215279\\StudyGo\\TeaCode\\day06\\04log_demo\\xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	log.SetOutput(fileObj)

	for {
		log.Println("这是一条测试的日志111")
		time.Sleep(time.Second * 3)
	}
}
