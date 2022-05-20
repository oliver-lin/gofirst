package main

import (
	"log"
)

func main() {
	log.Println("这是一条日志")
	// log.Fatalln("触发fatalln")
	// log.Panicln("触发panicln")
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条set flags 后的log")
}
