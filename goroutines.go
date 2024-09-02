package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int) {
	for i := 1; i < 20; i++ {
		ch <- i
	}
	close(ch)
}

func worker(id int, wg *sync.WaitGroup, in chan int, out chan int64) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	for value := range in {
		fmt.Printf("Worker %d received %d\n", id, value)
		var res int64 = 1
		for i := 2; i <= value; i++ {
			res *= int64(i)
		}
		out <- res
	}

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	ch := make(chan int)
	out := make(chan int64)
	var wg sync.WaitGroup
	go producer(ch)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg, ch, out)
	}

	go func() {
		for res := range out {
			fmt.Printf("Res = %d\n", res)
		}
	}()

	wg.Wait()
	close(out)
}
