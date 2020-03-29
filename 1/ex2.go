package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "STOP" {
			return
		} else {
			fmt.Println(">", scanner.Text())
		}
	}
}
