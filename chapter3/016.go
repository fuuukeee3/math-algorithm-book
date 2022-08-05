package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var a, b int
	for i := 0; i < n; i++ {
		var nn int
		fmt.Scan(&nn)

		if i == 0 {
			a = nn
		} else {
			b = nn
			a = gcd(a, b)
		}
	}
	fmt.Println(a)
}

func gcd(a, b int) int {
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

	return answer
}
