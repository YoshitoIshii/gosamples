package main

import (
	"fmt"
	"runtime"
)

func main() {
	FooFunc()
}

func FooFunc() {
	BarFunc()
}
func BarFunc() {
	depth := 0
	if pc, _, _, ok := runtime.Caller(depth); ok {
		fmt.Printf("%s\n",runtime.FuncForPC(pc).Name())
	}
}

