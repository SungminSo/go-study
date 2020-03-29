package main

import "fmt"

func main() {
	sliceA := make([]int, 0, 3)

	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))

	}
	fmt.Println(sliceA)
}
