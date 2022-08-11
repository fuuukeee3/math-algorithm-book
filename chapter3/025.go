package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var a []float64
	for i := 0; i < n; i++ {
		var n1 float64
		fmt.Scan(&n1)

		a = append(a, n1)
	}

	var b []float64
	for i := 0; i < n; i++ {
		var n2 float64
		fmt.Scan(&n2)

		b = append(b, n2)
	}

	var result float64
	for i := 0; i < n; i++ {
		result += (a[i] * 2.0 / 6.0) + (b[i] * 4.0 / 6)
	}
	fmt.Println(result)
}
