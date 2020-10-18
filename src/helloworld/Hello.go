package main

import "fmt"
import "./mathutil"

//常量的定义可以被用作枚举值
const (
	a = "abc"
	b = "hello"
)

func main() {
	fmt.Println("hello world")
	//这里引用一个定义包里面的方法，mathutil是用package来定义的，另外func的名字，首字母需要是大写，不然引用不到
	fmt.Println(mathutil.Add(1,2))

	test()

	fmt.Println(testReturnNums())

	testConst()

	testChannel()
}


func test() {
	// := 是一个申明语句 约等于 var x string = "hello world"
	var y string = "hello world"
	x := "hello world"
	println("x:" + x)
	println("y:" + y)
}

func testReturnNums()(int,int,string) {
	return 1,2,"3"
}

func testConst() {
	println(a,b)
}

func testChannel() {
	//创建一个channel 用于通信的管道
	messages:= make(chan string)
	go func() {	messages <- "hello"}()
	msg,ok := <- messages
	if(ok) {
		fmt.Println(msg)
	}
}
