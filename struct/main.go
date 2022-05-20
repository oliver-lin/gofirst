package main

import (
	"fmt"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type myStruct struct{ i int }

var v myStruct  // v是结构体类型变量
var p *myStruct // p是指向一个结构体类型变量的指针

func usct(mp *myStruct) {
	mp.i = 15
}

func main() {
	ms := new(struct1)

	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)

	v.i = 3
	fmt.Println(v.i)

	ms2 := new(myStruct)
	ms2.i = 1
	usct(ms2)
	fmt.Println(ms2.i)

	fmt.Println((*p).i)
}
