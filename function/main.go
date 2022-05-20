package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gofirst/function/closure"
)

var sum int = 0
var slice []int

func intSum(x ...int) {
	fmt.Printf("%T,%v \n", x, x)
	for _, v := range x {
		sum += v
	}
	fmt.Println(sum)
}

func Add(x, y int) int {
	return x + y
}
func Sub(x, y int) int {
	return x - y
}

type calculation func(int, int) int

var c calculation

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func dosth(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return Add, nil
	case "-":
		return Sub, nil
	default:
		return nil, errors.New("fail")
	}
}

// 闭包
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func incr() func() int {
	var x int
	return func() int {
		fmt.Println("初始x=", x)
		x++
		return x
	}
}

// 函数计数器
func counter(f func()) func() int {

	n := 0
	return func() int {

		f()
		n += 1
		fmt.Println("这时候的n=", n)
		return n
	}
}

// 测试的调用函数
func foo() {

	fmt.Println("call foo")
}

func main() {
	fmt.Println("func 的使用")
	/*

		intSum(1, 2, 3)

		// 函数类型
		c = Add
		fmt.Println(c(3, 5))

		// 函数作为参数
		r := calc(7, 3, Add)
		fmt.Println(r)

		// 函数作为返回值
		dfunc, _ := dosth("+")
		fmt.Println(dfunc(10, 20))

		//匿名函数
		add := func(x, y int) int {
			return x + y
		}
		fmt.Println(add(8, 9))

		// 闭包
		f := adder()
		fmt.Println(f(11))
		fmt.Println(f(11))

		name := makeSuffixFunc(".jpg")
		fmt.Println(name("test"))
	*/
	// var v_x int = 1

	// foo2 := closure.Foo2(v_x)
	// foo2()
	// fmt.Println(v_x)

	// foo1 := closure.Foo1(&v_x)
	// foo1()
	// fmt.Println(v_x)

	// closure.Foo3()

	closure.Foo4()

	// closure.Foo5()

	time.Sleep(time.Second)

	f := incr()
	println(f())
	println(f())

	// counter可以接受任何输入值和返回值为空的函数，同时返回一个闭包，在这里闭包的结果是函数的调用次数。
	// 本例子中，只用cnt对闭包的引用结束后，才会销毁闭包。一个闭包只有没有外界引用时，才会连同状态一起被销毁。
	// 如果记录需要参数的函数，需要单独给counter传递需要传递的参数类型。
	cnt := counter(foo)
	cnt()
	cnt()
	cnt()
	fmt.Println(cnt())

}
