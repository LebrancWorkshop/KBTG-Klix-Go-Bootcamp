package fizzbuzz

import (
	"testing"
)

type TestCase struct {
	name string
	input int
	want string
}

func TestFizzBuzz(t *testing.T) {
	testCases := []TestCase{
		{"1 should return 1", 1, "1"},
		{"2 should return 2", 2, "2"},
		{"3 should return 3", 3, "Fizz"},
		{"4 should return 4", 4, "4"},
		{"5 should return 5", 5, "Buzz"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := FizzBuzz(tt.input)

			if got != tt.want {
				t.Errorf("FizzBuzz(%v) should return %v, but got %v", tt.name, tt.want, got)
			}
		})
	}
}
