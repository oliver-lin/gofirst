package main

import (
	"fmt"
)

func main() {
	// a := 10
	// b := &a
	// fmt.Printf("a:%d ptr:%p\n", a, &a)
	// fmt.Printf("b:%p type:%T data:%v\n", b, b, b)
	// fmt.Println(&b)

	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	// 指针必须分配内存
	d := new(int)
	*d = 100
	fmt.Println(*d)

	e := make(map[string]int, 5)
	e["字段占位"] = 1
	fmt.Println(e, len(e))

}
