package main

import (
	"fmt"
)

func sqrt(x float64) (v float64) {
	v = x
	for i := 0; i < 10; i++ {
		v = v - (v*v-x)/(2*v)
	}
	return
}

func main() {
	fmt.Println(sqrt(2))
}
