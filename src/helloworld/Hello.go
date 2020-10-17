package main

import "fmt"
import "./mathutil"

func main() {
	fmt.Println("hello world")
	//这里引用一个定义包里面的方法，mathutil是用package来定义的，另外func的名字，首字母需要是大写，不然引用不到
	fmt.Println(mathutil.Add(1,2))
}
