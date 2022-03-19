package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAddition(t *testing.T) {
	result, _ := calculate(OperationAddition, 2, 2)
	var expected float64 = 4
	assert.Equal(t, expected, result)
}

func TestCalculateTable(t *testing.T) {
	tests := []struct {
		operand       string
		firstOperand  float64
		secondOperand float64
		expected      float64
	}{
		{
			operand:       OperationAddition,
			firstOperand:  2.0,
			secondOperand: 2.0,
			expected:      4.0,
		},
		{
			operand:       OperationMultiply,
			firstOperand:  4.0,
			secondOperand: 4.0,
			expected:      16.0,
		},
		{
			operand:       OperationDivide,
			firstOperand:  16.0,
			secondOperand: 2.0,
			expected:      8.0,
		},
	}

	for _, tc := range tests {
		got, _ := calculate(tc.operand, tc.firstOperand, tc.secondOperand)
		if !reflect.DeepEqual(tc.expected, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.operand, tc.expected, got)
		}
	}
}
