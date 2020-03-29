package main

import (
	"fmt"
)

func main() {
	// 전형적인 for문
	for i := 0; i < 100; i++ {
		if i % 20 == 0 {
			continue
		}

		if i == 95 {
			break
		}

		fmt.Print(i, " ")
	}

	// while문스러운 for문
	fmt.Println()
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}
	fmt.Println()

	// do...while문스러운 for문
	i = 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}

		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index:", i, "value: ", value)
	}
}
