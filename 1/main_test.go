package main

import (
	"testing"
)

func TestToString(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestToString",
			args: args{a: 42},
			want: "42",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toString(tt.args.a); got != tt.want {
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeString(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestMakeString",
			args: args{str: []string{"42", "052", "0x2A", "3.14", "Golang", "true", "1+2i"}},
			want: "42052x2A314Golangtrue1+2i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeString(tt.args.str...); got != tt.want {
				t.Errorf("makeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSHA256(t *testing.T) {
	type args struct {
		r []rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestHashSHA256",
			args: args{r: []rune("TEST")},
			want: "94ee059335e587e501cc4bf90613e0814f00a7b08bc7c648fd865a2af6a22cc2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashSHA256(tt.args.r); got != tt.want {
				t.Errorf("hashSHA256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddSalt(t *testing.T) {
	type args struct {
		r []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "TestAddSalt",
			args: args{r: []rune("TEST1TEST")},
			want: []rune("TESTgo-20241TEST"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addSalt(tt.args.r); string(got) != string(tt.want) {
				t.Errorf("addSalt() = %v, want %v", got, tt.want)
			}
		})
	}
}
