package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// randomGenerator returns a channel that emits n random integers
// between 0 and n. The channel is closed after n values have been sent.
func randomGenerator(n int) <-chan int {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			ch <- generator.Intn(n)
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	ch := randomGenerator(10)
	for num := range ch {
		fmt.Println(num)
	}

}
