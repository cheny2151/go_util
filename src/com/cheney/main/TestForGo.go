package main

import (
	"fmt"
	"sync"
)

var sum int
var	mutex = sync.Mutex{}

func main() {
	for i := 0; i <= 100; i++ {
		go Sum(i)
	}
	fmt.Print(sum)
}

func Sum(val int) {
	mutex.Lock()
	sum += val
	mutex.Unlock()
}
