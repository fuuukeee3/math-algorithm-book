package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	for i := 0; i < n; i++ {
		var nn int
		fmt.Scan(&nn)
		sum += nn
	}
	fmt.Println(sum)
}
