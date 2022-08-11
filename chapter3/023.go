package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var blue []float64
	for i := 0; i < n; i++ {
		var nn float64
		fmt.Scan(&nn)
		blue = append(blue, nn)
	}

	var red []float64
	for i := 0; i < n; i++ {
		var nn float64
		fmt.Scan(&nn)
		red = append(red, nn)
	}

	var blue_sum float64
	for _, b := range blue {
		blue_sum += b
	}

	var red_sum float64
	for _, b := range red {
		red_sum += b
	}

	fmt.Println((blue_sum / float64(n)) + (red_sum / float64(n)))
}
