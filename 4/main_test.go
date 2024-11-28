package main

import (
	"reflect"
	"testing"
)

func TestUniqueOfFirst(t *testing.T) {

	tests := []struct {
		name   string
		slice1 []string
		slice2 []string
		want   []string
	}{
		{"TestUniqueOfFirst1",
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{"banana", "date", "fig"},
			[]string{"apple", "cherry", "43", "lead", "gno1"}},

		{"TestUniqueOfFirst2",
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{"apple", "date", "fig", "date", "fig"},
			[]string{"banana", "cherry", "43", "lead", "gno1"}},

		{"TestUniqueOfFirst3",
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{}},

		{"TestUniqueOfFirst4",
			[]string{},
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{}},

		{"TestUniqueOfFirst5",
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{},
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := uniqueOfFirst(tc.slice1, tc.slice2); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("uniqueOfFirst() = %v, want %v", got, tc.want)
			}
		})
	}
}
