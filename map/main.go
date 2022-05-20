package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数  [0,n)
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	keys := make([]string, 0, 99)
	// values := make([]int, 0, 99)
	for k := range scoreMap {
		keys = append(keys, k)
	}
	// //对切片进行排序
	sort.Strings(keys)
	// //按照排序后的key遍历map
	for _, k := range keys {
		fmt.Println(k, scoreMap[k])
	}

	// ----
	// 元素为map类型的切片
	mapslice := make([]map[string]int, 3)
	for k, v := range mapslice {
		fmt.Printf("k type=%T, v type=%T, %v\n", k, v, mapslice[k])
	}
	// 对切片中的map元素进行初始化
	for i := 0; i < len(mapslice); i++ {
		height := 100
		age := 20
		high := 170
		mapslice[i] = map[string]int{"height": height + i, "age": age + i, "high": high + i}
	}
	fmt.Println(mapslice)

	// --- 值为切片类型的map
	mapslice2 := make(map[string][]string, 3)
	str := "abc"
	for _, v := range str {
		mapslice2[string(v)] = []string{"a", "b"}
	}
	fmt.Println(mapslice2)
}
