package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	// 创建trace
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// 启动trace goruntine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	// main
	fmt.Println("Hello GMP")
}
