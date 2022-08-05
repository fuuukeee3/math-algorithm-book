package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	var min int
	if a > b {
		min = b
	} else {
		min = a
	}

	answer := 1
	for i := 2; i <= min; i++ {
		if a%i == 0 && b%i == 0 {
			answer = i
		}
	}
	fmt.Println(answer)
}
