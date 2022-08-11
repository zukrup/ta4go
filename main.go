package main

import (
	"fmt"
	"time"
	bar "wlog3/ta4go/core/Bar"
)

func main() {
	fmt.Println("Ta4go challenge")
	b := bar.New(5.5, 5.5, 5.5, 5.5, 5.5, 5.5, 5, time.Now().Sub(time.Now()), time.Now(), time.Now())
	print(b)
}
