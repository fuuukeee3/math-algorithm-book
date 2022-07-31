package main

import (
	"fmt"
)

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	count := 0
	for red := 1; red <= n; red++ {
		for blue := 1; blue <= n; blue++ {
			if red+blue <= s {
				count++
			}
		}
	}
	fmt.Println(count)
}
