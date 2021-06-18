package service

import (
	"fmt"
	"sync"
)

func ChannelService() int {
	var wg sync.WaitGroup
	ch := make(chan int, 1)
	wg.Add(2)

	go func1(&wg, ch)
	go func2(&wg, ch)

	wg.Wait()
	val := <-ch
	return val
}

func func1(wg *sync.WaitGroup, ch chan int) {
	count := <-ch
	for i := 0; i < 5; i++ {
		count++
		fmt.Println("func1 >>>", count)
	}
	defer wg.Done()
	ch <- count
}

func func2(wg *sync.WaitGroup, ch chan int) {
	var count int = 0
	for i := 0; i < 5; i++ {
		count++
		fmt.Println(count)
	}
	defer wg.Done()
	ch <- count
}
