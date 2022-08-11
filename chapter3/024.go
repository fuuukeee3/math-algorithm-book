package main

import (
	"fmt"
)

type Answer struct {
	count int
	score int
}

func main() {
	var n int
	fmt.Scan(&n)

	var answers []Answer
	for i := 0; i < n; i++ {
		var n1, n2 int
		fmt.Scan(&n1, &n2)

		answer := Answer{
			count: n1,
			score: n2,
		}
		answers = append(answers, answer)
	}

	var result float64
	for _, answer := range answers {
		result += float64(answer.score) / float64(answer.count)
	}
	fmt.Println(result)
}
