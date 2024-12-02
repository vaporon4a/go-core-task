package main

import (
	"testing"
)

func TestRandomGenerator(t *testing.T) {
	tests := []struct {
		name string
		arg  int
	}{
		{
			name: "TestRandomGenerator10",
			arg:  10,
		},
		{
			name: "TestRandomGenerator100",
			arg:  100,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := make([]int, 0)
			for num := range randomGenerator(tc.arg) {
				res = append(res, num)
			}
			if len(res) != tc.arg {
				t.Errorf("randomGenerator must return %d values, but return %d", tc.arg, len(res))
			}
		})
	}

}
