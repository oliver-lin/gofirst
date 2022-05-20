package closure

import (
	"fmt"
)

// 指向整形数字的指针,数字+1，原值改变
func Foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Printf("foo1 val = %d\n", *x)
	}
}

// x 值作为永驻变量，长存 x+1, +1, +1
func Foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo2 val = %d\n", x)
	}
}

// 普通切片 range
func Foo3() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fmt.Printf("foo3 val = %d\n", val)
	}
}

func show(v interface{}) {
	fmt.Printf("foo4 val = %v\n", v)
}

// 启动了一个协程 非按 values 切片顺序，打印 foo4，需要执行 time.Sleep(time.Second)，停顿1
func Foo4() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go show(val)
	}
}

// 输出 foo5 val = 5，需要执行 time.Sleep(time.Second)
func Foo5() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go func() {
			fmt.Printf("foo5 val = %v\n", val)
		}()
	}
}

// 申请了goroutine, 并发？
var foo6Chan = make(chan int, 10)

func Foo6() {
	for val := range foo6Chan {
		go func() {
			fmt.Printf("foo6 val = %d\n", val)
		}()
	}
}

func Foo7(x int) []func() {
	var fs []func()
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fs = append(fs, func() {
			fmt.Printf("foo7 val = %d\n", x+val)
		})
	}
	return fs
}
