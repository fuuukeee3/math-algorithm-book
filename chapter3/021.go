package main

import (
	"fmt"
)

func factorical(num int) int {
	result := 1
	for i := num; i > 1; i-- {
		result *= i
	}
	return result
}

func main() {
	var n, r int
	fmt.Scan(&n, &r)

	fmt.Println(factorical(n) / (factorical(r) * factorical(n-r)))
}
