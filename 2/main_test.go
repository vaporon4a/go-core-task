package main

import (
	"reflect"
	"testing"
)

func TestAddElements(t *testing.T) {
	type args struct {
		slice []int
		elem  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "TestAddElements",
			args: args{slice: []int{1, 2, 3}, elem: 4},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addElements(tt.args.slice, tt.args.elem); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceExample(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "TestSliceExample",
			args: args{slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: []int{2, 4, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceExample(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceExample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "TestCopySlice",
			args: args{slice: []int{1, 2, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copySlice(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	type args struct {
		slice []int
		index int
	}
	type want struct {
		slice []int
		err   error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "TestRemoveElement",
			args: args{slice: []int{1, 2, 3, 4, 5}, index: 2},
			want: want{slice: []int{1, 2, 4, 5}, err: nil},
		},
		{
			name: "TestRemoveElementOutOfRange",
			args: args{slice: []int{1, 2, 3, 4, 5}, index: 7},
			want: want{slice: nil, err: errIndexOutOfRange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := removeElement(tt.args.slice, tt.args.index); !reflect.DeepEqual(got, tt.want.slice) || !reflect.DeepEqual(err, tt.want.err) {
				t.Errorf("removeElement() = %v, want %v", got, tt.want.slice)
			}
		})
	}
}
