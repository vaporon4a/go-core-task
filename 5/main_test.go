package main

import (
	"reflect"
	"testing"
)

func TestSlicesCrossing(t *testing.T) {

	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		want   bool
		result []int
	}{
		{
			name:   "TestSlicesCrossing1",
			slice1: []int{65, 3, 58, 678, 64},
			slice2: []int{64, 2, 3, 43},
			want:   true,
			result: []int{64, 3},
		},
		{
			name:   "TestSlicesCrossing2",
			slice1: []int{1, 1, 1, 678, 64},
			slice2: []int{64, 2, 64, 43},
			want:   true,
			result: []int{64},
		},
		{
			name:   "TestSlicesCrossing3",
			slice1: []int{},
			slice2: []int{64, 2, 64, 43},
			want:   false,
			result: []int{},
		},
		{
			name:   "TestSlicesCrossing4",
			slice1: []int{1, 1, 1, 678, 64},
			slice2: []int{},
			want:   false,
			result: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if flag, result := slicesCrossing(tc.slice1, tc.slice2); flag != tc.want || !reflect.DeepEqual(result, tc.result) {
				t.Errorf("slicesCrossing(%v, %v) = (%v, %v), want (%v, %v)", tc.slice1, tc.slice2, flag, result, tc.want, tc.result)

			}
		})
	}
}

func TestMax(t *testing.T) {

	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"TestMax1", 1, 2, 2},
		{"TestMax2", 2, 1, 2},
		{"TestMax3", 2, 2, 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := max(tc.a, tc.b); got != tc.want {
				t.Errorf("max(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
