package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Calculate(x int) (result int) {
	result = x * 6
	return result
}

func Calculate1(a int) (result int) {
	result = a - 1
	return result
}
func TestCalculate3(t *testing.T) {
	assert.Equal(t, Calculate1(2), 1)
}
