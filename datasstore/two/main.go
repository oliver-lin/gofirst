package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// GOPATH：
//  /Users/linning/go
// 当前项目目录：
//  /Users/linning/go/src/github.com/backend/gofirst
// 当前文件目录：
//  /Users/linning/go/src/github.com/backend/gofirst/datasstore/two/main.go
// 当前文件夹目录：
//  /Users/linning/go/src/github.com/backend/gofirst/datasstore/two
func Dirs() map[string]string {
	dirs := make(map[string]string, 3)
	projpath, _ := os.Getwd()
	dirs["proj_path"] = projpath
	_, fullFilename, _, _ := runtime.Caller(0)
	dirs["file_path"] = fullFilename
	basepath := filepath.Dir(fullFilename)
	dirs["curr_dir"] = basepath
	return dirs
}

func main() {
	dirs := Dirs()

	data := []byte("Hello world!\n")
	file1 := dirs["curr_dir"] + "/" + "datafile1"
	err := ioutil.WriteFile(file1, data, 0644)
	if err != nil {
		panic(err)
	}
	read, _ := ioutil.ReadFile(file1)
	fmt.Printf("%+v\n", string(read))

	// 文件写入+读取方式二
	file2 := dirs["curr_dir"] + "/" + "datafile2"
	fh1, _ := os.Create(file2)
	defer fh1.Close()

	bytes, _ := fh1.Write(data)
	fmt.Printf("写入了：%+v byte\n", bytes)

	fh2, _ := os.Open(file2)
	defer fh2.Close()

	readout := make([]byte, len(data))
	bytecnt, _ := fh2.Read(readout)
	fmt.Printf("%+v,%+v\n", bytecnt, string(readout))

	file3 := dirs["curr_dir"] + "/" + "datafile3.txt"
	//生成[15，88]之间的随机数,括号左包含右不包含
	// 67868265
	str := ""
	start := time.Now().UnixMicro()
	randnumbers2 := map[int]int{}
	randnumbers3 := map[int]int{}
	randnumbers1 := []int{}
	randnumbers := make(map[int]int, 100000)

	// for i := 0; i < 100000; i++ {
	i := 0
	for {

		if i == 100000 {
			break
		}
		rand.Seed(time.Now().UnixNano())
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rnd := r.Intn(10000000) + 60000000 //(88-15 )+15
		if _, ok := randnumbers[rnd]; ok {
			continue
		}
		randnumbers[rnd] = 1
		randnumbers1 = append(randnumbers1, rnd)
		numstr := strconv.Itoa(rnd) + "\n"
		str += numstr
		// fmt.Printf("rand is %v\n", rnd)
		i++
	}
	drua := time.Now().UnixMicro() - start
	str = strings.TrimSuffix(str, "\n")
	err = ioutil.WriteFile(file3, []byte(str), 0644)
	if err != nil {
		panic(err)
	}
	fmtdra := float64(drua) / float64(1e6)
	fmt.Println(fmtdra)

	for _, v := range randnumbers1 {
		if _, ok := randnumbers2[v]; ok {
			randnumbers3[v] = randnumbers2[v] + 1
		} else {
			randnumbers2[v] = 1
		}
	}
	fmt.Println(randnumbers3)

}
