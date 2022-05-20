package main

import (
	"fmt"
	"unicode/utf8"
)

var (
	v_u8      uint8  = 254        // 最大255，有符号则对半分 -128~127
	v_u16     uint16 = 65534      // 65535
	v_u32     uint32 = 4290000000 // 42亿
	v_rune           = "白菜萝卜"
	v_chinese        = "我是中国人 I am chinese"
)

func main() {
	fmt.Println("类型")
	fmt.Printf("%d,%d,%d \n", v_u8, v_u16, v_u32)

	vvrune := []rune(v_rune)
	vvrune[0] = '红'
	fmt.Println(len(vvrune), string(vvrune))

	fmt.Println("origin len=", len(v_chinese))
	fmt.Println("rune len=", len([]rune(v_chinese)))
	fmt.Println(utf8.RuneCountInString(v_chinese))

	for i, v := range []rune(v_chinese) {
		fmt.Printf("i=%d,v=%s\n", i, string(v))
	}

	var v_str1 = "小手25是什么"
	s1 := []rune(v_str1)
	fmt.Println(string(s1[:4]))

	// 在N个字符串找到找没有重复字符，且字符串总长度最长的那个
	v_str2 := []string{
		"abccefg",
		"bbbb",
		"cdcdccdd",
		"ddddddddddd",
		"来个长点的字符串，微信号5",
	}
	idx := 0
	maxlen := 0
	for i := 0; i < len(v_str2); i++ {
		var repeat bool
		tmpArr := map[int32]int{}
		for k, v := range []rune(v_str2[i]) {
			_, ok := tmpArr[v]
			if ok {
				repeat = true
				break
			}
			// if tmpArr[v] != 0 && len(tmpArr) > 0 {
			// 	repeat = true
			// 	break
			// }
			tmpArr[v] = k
		}
		// 不重复，且对比上一个
		if !repeat && len(tmpArr) > maxlen {
			maxlen = len(tmpArr)
			idx = i
		}
	}
	fmt.Println("idx:", idx, "element:", v_str2[idx], "maxlen:", maxlen)

	// map 类型
	mapscore := make(map[string]int, 5)
	mapscore["张三"] = 90
	mapscore["lisi"] = 85
	mapscore["wangwu"] = 87
	fmt.Printf("%T", mapscore)
}
