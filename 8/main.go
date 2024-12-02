package main

import (
	"fmt"
	"time"
)

type customWG struct {
	semaphore chan struct{}
	counter   int
}

// NewCustomWG creates a new custom WaitGroup with a capacity of n.
//
// The custom WaitGroup is implemented as a channel semaphore. The Add
// method adds n to the counter and sends n empty structs on the channel.
// The Wait method waits until the counter is 0 and then closes the channel.
// The Done method decrements the counter and receives one empty struct from
// the channel.
func NewCustomWG(n int) *customWG {
	return &customWG{
		semaphore: make(chan struct{}, n),
	}
}

// Add increments the counter by n and sends n empty structs on the semaphore channel.
//
// This method is used to indicate that n goroutines are to be added to the custom wait group.
// It should be called before starting the goroutines that need to be tracked.
func (c *customWG) Add(n int) {
	c.counter += n
	for i := 0; i < n; i++ {
		c.semaphore <- struct{}{}
	}
}

// Done decrements the counter by 1 and receives an empty struct from the semaphore channel.
//
// This method is used to indicate that a goroutine has completed its task.
// It should be called once for each goroutine added with the Add method.
func (c *customWG) Done() {
	if c.counter > 0 {
		c.counter--
		<-c.semaphore
	}
}

// Wait blocks until the counter is 0. It is not a traditional wait group
// implementation, as it does not use a condition variable or a channel to
// signal when the counter is 0. Instead, it periodically checks the counter
// and returns when it is 0. This is a simple but not very efficient
// implementation, but it is good enough for our use case.
func (c *customWG) Wait() {
	for c.counter > 0 {
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	const numbers = 10
	wg := NewCustomWG(numbers)
	wg.Add(numbers)
	for i := 0; i < numbers; i++ {
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second * 5)
			fmt.Printf("Goroutine %d is finished.\n", i)
		}(i)
	}

	wg.Wait()
	fmt.Println("All jobs is finished.")
}
