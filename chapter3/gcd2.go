package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	answer := 1
	for {
		if a < b {
			b = b % a
		} else {
			a = a % b
		}

		if a == 0 || b == 0 {
			if a > b {
				answer = a
			} else {
				answer = b
			}
			break
		}
	}

	fmt.Println(answer)
}
