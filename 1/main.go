package main

import (
	"fmt"
	"awesomeProject/pkg/stringutil"
	"github.com/mactsouk/go/simpleGitHub"
)

func main() {
	fmt.Println("Hello world!")
	val := stringutil.Reverse("ABCDE")
	fmt.Println(val)

	fmt.Println(simpleGitHub.AddTwo(5, 6))
}