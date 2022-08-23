package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	var h []int
	for i := 0; i < n; i++ {
		var hh int
		fmt.Scan(&hh)

		h = append(h, hh)
	}

	var dp []int
	for i := 0; i < n; i++ {
		if i == 0 {
			dp = append(dp, 0)
		}

		if i == 1 {
			dp = append(dp, int(math.Abs(float64(h[i])-float64(h[i-1]))))
		}

		if i >= 2 {
			i1 := dp[i-1] + int(math.Abs(float64(h[i-1])-float64(h[i])))
			i2 := dp[i-2] + int(math.Abs(float64(h[i-2])-float64(h[i])))
			if i1 >= i2 {
				dp = append(dp, i2)
			} else {
				dp = append(dp, i1)
			}
		}
	}
	fmt.Println(dp[n-1])
}
