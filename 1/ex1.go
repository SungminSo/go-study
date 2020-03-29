package main

import (
	"os"
	"fmt"
	"strconv"
	"errors"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please one or more float!")
		os.Exit(1)
	}

	arguments := os.Args
	var err error = errors.New("An error")
	k := 1
	sum := 0.

	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of the arguments is a float!")
			return
		}
		_, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}

	for i := 1; i < len(arguments); i++ {
		 n, _ := strconv.ParseFloat(arguments[i], 64)
		 sum += n
	}

	fmt.Println("Sum:", sum)
	fmt.Println("Mean:", sum / float64(len(arguments)))
}
