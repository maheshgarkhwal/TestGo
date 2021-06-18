package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	// assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 12},
		{-1, -6},
		{0, 0},
	}

	for _, test := range tests {
		assert.Equal(t, Calculate(test.input), test.expected)
	}
}

func TestCalculate2(t *testing.T) {
	assert.Equal(t, Calculate1(2), 1)
}
