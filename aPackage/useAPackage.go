package main

import (
	aPackage2 "awesomeProject/src/aPackage"
	"fmt"
)

func main() {
	fmt.Println("Using aPackage!")
	aPackage2.A()
	aPackage2.B()
	fmt.Println(aPackage2.MyConstant)
}