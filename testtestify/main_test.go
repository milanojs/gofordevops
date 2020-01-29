package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

/*
func TestCalculate(t *testing.T) {
	// classic method to test a function without assert
	/*
		if Calculate(2) != 4 {
			t.Error("Expected 2 + 2 to equal 4")
		}
*/
/*
	assert.Equal(t, Calculate(2), 4)

}*/
func TestCalculate(t *testing.T) {

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected)
	}
}
