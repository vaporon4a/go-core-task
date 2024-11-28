package main

import (
	"reflect"
	"testing"
)

func InitMap() *StringIntMap {
	m := NewHashMap()
	m.Add("test", 42)
	m.Add("TEST2", 56)
	m.Add("TEST3", 78)
	return m
}

func TestAdd(t *testing.T) {

	tests := []struct {
		name  string
		key   string
		value int
		want  int
	}{
		{"TestAddElementsOK", "someelement", 66, 4},
		{"TestAddElementsExist", "test", 99, 4},
		{"TestAddElementsTwo", "ANOTHERELEMENT2", 78, 5},
	}

	m := InitMap()

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			m.Add(tc.key, tc.value)

			if got := m.size; got != tc.want {
				t.Errorf("size = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type wants struct {
		bucket *bucketNode
		size   int
	}

	tests := []struct {
		name  string
		key   string
		wants wants
	}{
		{"TestRemoveElementsOk", "test", wants{nil, 2}},
		{"TestRemoveElementsNotExist", "test", wants{nil, 2}},
	}

	m := InitMap()

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			m.Remove(tc.key)

			if got := m.array[hash(tc.key, m.arraySize)].head; !reflect.DeepEqual(got, tc.wants.bucket) {
				t.Errorf("bucket = %v, want %v", got, tc.wants.bucket)
			}
			if got := m.size; got != tc.wants.size {
				t.Errorf("size = %v, want %v", got, tc.wants.size)
			}
		})
	}
}

func TestCopy(t *testing.T) {

	m := InitMap()

	tests := []struct {
		name string
		want *StringIntMap
	}{
		{"TestCopy", m},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := m.Copy(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Copy() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestExists(t *testing.T) {
	m := InitMap()

	tests := []struct {
		name string
		key  string
		want bool
	}{
		{"TestExistsOK", "test", true},
		{"TestExistsNotExist", "test2", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := m.Exist(tc.key); got != tc.want {
				t.Errorf("Exist() = %v, want %v", got, tc.want)
			}
		})
	}

}

func TestGet(t *testing.T) {
	m := InitMap()

	tests := []struct {
		name string
		key  string
		want int
		flag bool
	}{
		{"TestGetOk", "test", 42, true},
		{"TestGetNotExist", "test2", 0, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got, flag := m.Get(tc.key); got != tc.want || flag != tc.flag {
				t.Errorf("Get(%s) = (%v, %v), want (%v, %v)", tc.key, got, flag, tc.want, tc.flag)
			}
		})
	}
}

func TestHash(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		arraySize int
		want      int
	}{
		{"TestHash", "test", 10, 8},
		{"TestHash2", "test2", 5, 3},
		{"TestHash3", "test3", 10, 9},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := hash(tc.key, tc.arraySize); got != tc.want {
				t.Errorf("hash() = %v, want %v", got, tc.want)
			}
		})
	}
}
