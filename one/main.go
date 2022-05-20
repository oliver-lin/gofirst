package main // 声明 main 包，表明当前是一个可执行程序

import (
	"fmt" // 导入内置 fmt 包
	"strings"
)

var (
	v1_str        = "类型推导"
	v2_str string = "指定类型"
	v3_int int
)

const (
	a = 33 // 下一行没赋值，直接默认使用第一行
	b
	i = 35
	j
	c, d = iota + 1, iota + 2 // 第一行 iota = 0
	e, f                      // 第二行 iota = 1，默认使用第一行的表达式 e := 1+1, f := 1+2
	g, h                      // 第二行 iota = 2，默认使用第一行的表达式 e := 2+1, f := 2+2
)

type Student struct {
	Name string
	Age  int
}

func main() { // main函数，是程序执行的入口
	fmt.Println("Hello World!") // 在终端打印 Hello World!
	// var声明后直接 = ，如果没声明，则 := 赋值
	v3_int = 1
	v4_float := 0.66
	fmt.Println(v1_str, v2_str, v3_int, v4_float)

	fmt.Println(a, b, c, d, e, f, g, h, i, j)

	fmt.Println("%v,%+v,%#v 的区别")
	stu := Student{Name: "xiaoming", Age: 20}
	fmt.Printf("只输出值---%v\n", stu)
	fmt.Printf("字段名+值---%+v\n", stu)
	fmt.Printf("结构体+字段名+值---%#v", stu)

	hw := "hello world"
	hw = strings.Trim(hw, "hrd")
	fmt.Printf("%+v", hw)

}
