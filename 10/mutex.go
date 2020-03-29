package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	m sync.Mutex
	v1 int
)

func change(i int) {
	m.Lock()
	fmt.Println("Lock", i)
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1%10 == 0 {
		v1 = v1 - 10*i
	}
	m.Unlock()
	fmt.Println("Unlock", i)
}

func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	numGR, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup

	fmt.Println("numGR:", numGR)
	fmt.Printf("%d ", read())
	for i := 0; i < numGR; i++ {
		waitGroup.Add(1)
		fmt.Println("asdfasf")
		go func(i int) {
			defer waitGroup.Done()
			change(i)
			fmt.Printf("-> %d", read())
		}(i)
	}

	 waitGroup.Wait()
	fmt.Printf("-> %d\n", read())
}