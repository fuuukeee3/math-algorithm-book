package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var results []int
	for i := 2; i <= n; i++ {
		flg := true
		for j := 2; j <= i-1; j++ {
			if i%j == 0 {
				flg = false
				break
			}
		}
		if flg {
			results = append(results, i)
		}
	}

	for i, result := range results {
		fmt.Print(result)
		if i != len(results) {
			fmt.Print(" ")
		} else {
			fmt.Println("")
		}
	}
}
