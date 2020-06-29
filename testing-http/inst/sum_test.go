package sum_test

import (
	"testing"
)

func TestInts(t *testing.T) {

	tt := []struct {
		name    string
		numbers []int
		sum     int
		message string
	}{
		{"sum 1 to 5", []int{1, 2, 3, 4, 5}, 15, "The sum to 1 to 5 should be %v ; got %v"},
		{"sum 0", nil, 0, "The sum of 0 should should be %v ; got %v"},
		{"sum 1+-1", []int{1, -1}, 0, "The sum to 1 plus -1 should be %v ; got %v"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := Ints(tc.numbers...)
			if s != tc.sum {
				t.Errorf(tc.message, tc.sum, s)

			}
		})
	}
}
func TestFoo(t *testing.T) {
	t.Fatal("this should always fail")
}
