package main

import (
	"fmt"
	"sync"
)

func main() {
	n := [][]int{{2, 6, 9, 24}, {7, 3, 94, 3, 0}, {4, 2, 8, 35}}
	var wg sync.WaitGroup

	wg.Add(len(n))

	for i := 0; i < len(n); i++ {
		go func(i int) {
			defer wg.Done()
			sum(n[i], i)
		}(i)
	}
	wg.Wait()
}

func sum(arr []int, i int) {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Printf("slice %d: %d\n", i+1, sum)
}
