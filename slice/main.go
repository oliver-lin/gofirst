package main

import (
	"fmt"
)

var (
	slice_int  []int
	slice_bool []bool
)

func main() {
	slice := []string{"a", "b", "c", "d", "e"}
	newslice := slice[1:3]
	for i, v := range newslice {
		fmt.Println(i, v)
	}
	slice_int = []int{1, 2}
	slice_bool = []bool{true, false, true}
	fmt.Println(slice_int, slice_bool)

	// 删除元素d（实则：追加切片覆盖，只是把e覆盖到d上）
	newslice2 := append(slice[:3], slice[4:]...)
	fmt.Println("取删除元素d后的新切片：", newslice2, "底层数组容量不变，仍然是：", cap(newslice2), "现在的底层数组：", slice)
	fmt.Printf("现阶段slice地址：%p\nnewslice2地址：%p \n", slice, newslice2)
	newaddrslice := make([]string, 4)
	copy(newaddrslice, newslice2)
	fmt.Printf("地址：%p，值是：%v，容量：%d \n", newaddrslice, newaddrslice, cap(newaddrslice))
	newaddrslice2 := append(newaddrslice, "h", "i", "j")
	fmt.Printf("地址：%p，值是：%v，容量：%d \n", newaddrslice2, newaddrslice2, cap(newaddrslice2))
	fmt.Println("我是分割线---------------------")

	newslice3 := append(slice[:3], slice[3:]...)
	fmt.Println("覆盖slice元素:", newslice3, cap(newslice3))

	slice2 := [...][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Printf("len=%d,type=%T,value=%v \n", len(slice2), slice2, slice2)

}
