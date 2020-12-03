package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}
func main() {
	done := make(chan bool, 1)
	//使用go关键字 开启channel
	fmt.Println("start go")
	go worker(done)

	//接收done channel中的消息，等待任务完成
	<-done

	fmt.Println("worker done")
}

