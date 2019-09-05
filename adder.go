package main

import (
	"fmt"
	"github.com/tisnik/intvalues"
)

func Add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(Add(intvalues.Zero, intvalues.Zero))
	fmt.Println(Add(intvalues.One, 0))
	fmt.Println(Add(1, 2))
	fmt.Println(Add(-1, 2))
}
