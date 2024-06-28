package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("1 + 2 = 3", func(t *testing.T) {
		got := Sum(1, 2)
		want := 3

		if got != want {
			t.Errorf("Test Failed: Got: %v, but Want: %v", got, want)
		}
	})
}
