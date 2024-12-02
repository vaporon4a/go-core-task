package main

import (
	"sync"
	"testing"
	"time"
)

func TestNewCustomWG(t *testing.T) {
	t.Run("TestNewCustomWG", func(t *testing.T) {
		wg := NewCustomWG(10)
		if wg.counter != 0 {
			t.Errorf("expected counter to be 0, got %d", wg.counter)
		}
		if cap(wg.semaphore) != 10 {
			t.Errorf("NewCustomWG channel capacity is bad.")
		}
	})
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "TestAdd-10",
			args: 10,
			want: 10,
		},
		{
			name: "TestAdd-100",
			args: 100,
			want: 100,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			wg := NewCustomWG(tc.args)
			wg.Add(tc.args)
			if wg.counter != tc.want || cap(wg.semaphore) != tc.want {
				t.Errorf("counter or chan capacity wrong, got = %d; want = %d", wg.counter, tc.want)
			}
		})
	}
}

func TestDone(t *testing.T) {
	name := "TestDone"
	want := 9
	t.Run(name, func(t *testing.T) {
		wg := NewCustomWG(10)
		wg.Add(10)
		wg.Done()
		if wg.counter != want || len(wg.semaphore) != want {
			t.Errorf("Done() is not working, counter=%d; len of semaphore=%d; want=%d", wg.counter, len(wg.semaphore), want)
		}
	})
}

func TestWait(t *testing.T) {
	const numbers = 10
	wg := NewCustomWG(numbers)

	var wgSync sync.WaitGroup
	wgSync.Add(numbers)

	for i := 0; i < numbers; i++ {
		go func(i int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			wgSync.Done()
		}(i)
	}

	wgSync.Wait()
	wg.Wait()

	if wg.counter != 0 {
		t.Errorf("expected counter to be 0 after all Done calls, got %d", wg.counter)
	}

}
