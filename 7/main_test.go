package main

import (
	"testing"
)

func TestMergeChannels(t *testing.T) {
	tests := []struct {
		name string
		args []int
	}{
		{
			name: "TestMergeChannels6",
			args: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "TestMergeChannels30",
			args: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a := make(chan int)
			b := make(chan int)
			c := make(chan int)

			go func() {
				for i := 0; i < len(tc.args)/3; i++ {
					a <- tc.args[i]
				}
				close(a)
			}()

			go func() {
				for i := len(tc.args) / 3; i < (len(tc.args)/3)*2; i++ {
					b <- tc.args[i]
				}
				close(b)
			}()

			go func() {
				for i := (len(tc.args) / 3) * 2; i < len(tc.args); i++ {
					c <- tc.args[i]
				}
				close(c)
			}()
			res := make([]int, 0)
			for num := range mergeChannels(a, b, c) {
				res = append(res, num)
			}

			if len(res) != len(tc.args) {
				t.Errorf("mergeChannels must return %d values, but return %d", len(tc.args), len(res))
			}
		})
	}
}
