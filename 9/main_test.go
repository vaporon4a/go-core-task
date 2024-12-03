package main

import "testing"

func TestTripple(t *testing.T) {
	tests := []struct {
		name string
		args uint8
		want float64
	}{
		{
			name: "TestTripple1",
			args: 1,
			want: 1,
		},
		{
			name: "TestTripple3",
			args: 3,
			want: 27,
		},
		{
			name: "TestTripple5",
			args: 5,
			want: 125,
		},
		{
			name: "TestTripple10",
			args: 10,
			want: 1000,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tripple(tc.args)
			if result != tc.want {
				t.Errorf("Result of tripple(n uint8) not OK; result=%f; want=%f;", result, tc.want)
			}
		})
	}
}

func TestNumberConveer(t *testing.T) {
	chuint8 := make(chan uint8)
	chfloat64 := make(chan float64)
	wants := []float64{1, 8, 27, 64, 125, 216, 343, 512, 729, 1000}

	go func() {
		defer close(chuint8)
		for i := 1; i <= 10; i++ {
			chuint8 <- uint8(i)
		}
	}()

	go numberConveer(chuint8, chfloat64)

	i := 0

	for result := range chfloat64 {
		if result != wants[i] {
			t.Errorf("Expected %f but got %f", wants[i], result)
		}
		i++
	}
}
