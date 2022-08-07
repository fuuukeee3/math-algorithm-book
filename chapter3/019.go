package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	counts := map[string]int{
		"red":    0,
		"yellow": 0,
		"blue":   0,
	}

	for i := 0; i < n; i++ {
		var nn int
		fmt.Scan(&nn)

		switch nn {
		case 1:
			counts["red"]++
		case 2:
			counts["yellow"]++
		case 3:
			counts["blue"]++
		}
	}

	red := counts["red"] * (counts["red"] - 1) / 2
	yellow := counts["yellow"] * (counts["yellow"] - 1) / 2
	blue := counts["blue"] * (counts["blue"] - 1) / 2

	fmt.Println(red + yellow + blue)

}
