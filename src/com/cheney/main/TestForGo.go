package main

import (
	"com/cheney/untils"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var sum int
var mutex = sync.Mutex{}

func main() {
	/*for i := 0; i <= 100; i++ {
		go Sum(i)
	}
	fmt.Println(sum)*/
	file := "/Users/chenyi/yy/translation/lang.xlsx"
	excel, _ := untils.ReadExcel(file)
	cell := excel.ReadAllCell(0)
	results := make([]string, 1)
	for _, row := range cell {
		for i, col := range row {
			if i == 2 || i == 5 || i == 8 {
				if col != "" {
					results = append(results, col)
				}
			}
		}
	}
	join := strings.Join(results, "\",\"")
	print(join)
}

func Sum(val int) {
	mutex.Lock()
	sum += val
	println(GetCurrentThreadId())
	mutex.Unlock()
}

func GetCurrentThreadId() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
