package main

import (
	"fmt"
)

func main() {
	aMap := map[string]int{}
	aMap["test"] = 1

	theMap := map[string]int{}
	theMap = nil
	fmt.Println(theMap)
	theMap["test"] = 1
}
