package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	counts := map[int]int{}
	for i := 0; i < n; i++ {
		var nn int
		fmt.Scan(&nn)

		counts[nn]++
	}

	result := counts[100]*counts[400] + counts[200]*counts[300]
	fmt.Println(result)
}
