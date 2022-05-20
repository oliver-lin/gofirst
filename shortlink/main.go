package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

var CHARS = "InsV3Sf0obzp2i4gj1yYGqQv6wUtmBxlMAP7KHd8uTXFk9aRJWNC5EOhZDcLer"

const (
	// 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
	SCALE = 62
	REGEX = "^[0-9a-zA-Z]+$"
	NUM   = 6
)

func RandomStr(str string) string {
	chars := []rune(str)
	for i := len(chars) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		chars[i], chars[num] = chars[num], chars[i]
	}
	return string(chars)
}

func Encode10To62(val int) string {
	if val < 0 {
		panic("val cannot be negative.")
	}
	str := ""
	var remainder int
	for math.Abs(float64(val)) > SCALE-1 {
		remainder = int(val % SCALE)
		str = string(CHARS[remainder]) + str
		val = val / SCALE
	}
	str = string(CHARS[val]) + str
	//for i := len(str); i < NUM; i++ {//    str = string(CHARS[0]) + str
	//}
	return str
}

func Decode62To10(val string) uint {
	if match, _ := regexp.MatchString(REGEX, val); !match {
		panic("input illegal.")
	}
	var result uint = 0
	index, length := 0, len(val)
	for i := 0; i < length; i++ {
		index = strings.Index(CHARS, string(val[i]))
		result += uint(index * int(math.Pow(float64(SCALE), float64(length-i-1))))
	}
	return result
}

var (
	// 64进制使用到的字符列表(编码使用)
	endCode = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/")
	// 64进制使用到的字符map(解码使用)
	deCode = map[rune]int{}

	// 64进制
	SYSTEM uint64 = 64
)

func init() {
	for k, v := range endCode {
		deCode[v] = k
	}
}

// 编码
func Encode(id uint64) (string, error) {
	var data []rune
	for {
		var r rune   // 下标指向的字符
		var k uint64 // 64进制字符数组下标
		if id < SYSTEM {
			k = id
			r = endCode[k]
			data = append([]rune{r}, data...)
			break
		} else {
			k = id % SYSTEM
			r = endCode[k]
			data = append([]rune{r}, data...)

			id = (id - k) / SYSTEM
		}
	}

	return string(data), nil
}

// 解码
func Decode(str string) (uint64, error) {
	strRune := []rune(str) // 字符串转字符数组

	l := len(strRune)
	zs := l - 1 // 当前位指数
	var value uint64
	for i := 0; i < l; i++ {
		number, err := searchV(strRune[i])
		if err != nil {
			return 0, err
		}

		value += uint64(math.Pow(float64(SYSTEM), float64(zs))) * number
		zs--
	}

	return value, nil
}

// 过去字符在定义好的字符数组中的位置
func searchV(rune rune) (uint64, error) {

	k, ok := deCode[rune]
	if !ok {
		return 0, errors.New("字符不存在")
	}

	return uint64(k), nil

}

// ------------------------------------------------------------

var originlink string = "https://studygolang.com/topics/15244#reply1"

func main() {
	// rdstr := RandomStr(originlink)
	// shortChar := Encode10To62(18123)
	// id := Decode62To10(shortChar)
	// fmt.Println(rdstr, shortChar, id, Decode62To10(rdstr))
	// fmt.Println("--------")

	for _, e := range os.Environ() {
		fmt.Println(e)
	}

}
