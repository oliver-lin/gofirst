package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

// 年龄会增加，因为使用了指针，引用类型
func (p *person) guonian1() {
	p.age++
}

// 年龄不会变化，方法值传递 copy , paste
func (p person) guonian0() {
	p.age++
}

// New... 开头的构造函数名 (初始化变量)，结构名 { return 初始化结构体变量 }
func Newperson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func main() {
	fmt.Println("值接收和引用")
	p := Newperson("linning", 26)
	p.guonian0()
	fmt.Println(p.name, "白过年了，年龄不变", p.age)
	p.guonian1()
	fmt.Println(p.name, "过年了，年龄增长一岁", p.age)
}
