package main

import (
	"fmt"
	"sync"
)

func mergeChannels(chs ...<-chan int) <-chan int {
	mergegCh := make(chan int)
	wg := &sync.WaitGroup{}
	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			for num := range ch {
				mergegCh <- num
			}
		}(ch, wg)
	}

	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(mergegCh)
	}(wg)

	return mergegCh
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{20, 10, 30} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{300, 200, 100} {
			c <- num
		}
		close(c)
	}()

	for num := range mergeChannels(a, b, c) {
		fmt.Println(num)
	}
}
