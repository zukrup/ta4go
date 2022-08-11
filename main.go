package main

import (
	"fmt"
	bar "wlog3/ta4go/core/Bar"
)

func main() {
	fmt.Println("Ta4go challenge")
	b := bar.NewFakeBar()
	fmt.Println(b.ToString())
}
