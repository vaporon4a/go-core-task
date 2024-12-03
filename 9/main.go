package main

import "fmt"

// numberConveer takes a channel of uint8 and a channel of float64, and it
// closes the float64 channel after writing the cubed value of each number
// from the uint8 channel to the float64 channel.
func numberConveer(chuint8 <-chan uint8, chfloat64 chan<- float64) {
	defer close(chfloat64)
	for num := range chuint8 {
		chfloat64 <- tripple(num)
	}
}

// tripple takes a uint8 number and returns its cube as a float64.
func tripple(n uint8) float64 {
	return float64(n) * float64(n) * float64(n)
}

func main() {
	chuint8 := make(chan uint8)
	chfloat64 := make(chan float64)

	go func(chuint8 chan<- uint8) {
		defer close(chuint8)
		for i := 1; i <= 10; i++ {
			chuint8 <- uint8(i)
		}
	}(chuint8)

	go numberConveer(chuint8, chfloat64)

	for num := range chfloat64 {
		fmt.Println(num)
	}

}
