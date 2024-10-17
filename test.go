package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		target   int
		expected [][]int
	}{
		{
			name:     "Valid pairs with duplicates",
			numbers:  []int{1, 1, 2, 3, 7},
			target:   4,
			expected: [][]int{{1, 3}, {0, 3}},
		},
		{
			name:     "No pairs found",
			numbers:  []int{1, 2, 5},
			target:   10,
			expected: [][]int{},
		},
		{
			name:     "Valid pairs with same number",
			numbers:  []int{1, 2, 3, 3, 5},
			target:   6,
			expected: [][]int{{2, 3}, {0, 4}},
		},
		{
			name:     "Invalid input (empty array)",
			numbers:  []int{},
			target:   4,
			expected: [][]int{},
		},
		{
			name:     "Target is negative",
			numbers:  []int{-1, -2, -3, -4},
			target:   -5,
			expected: [][]int{{2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TwoSum(tt.numbers, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TwoSum(%v, %d) = %v; want %v", tt.numbers, tt.target, result, tt.expected)
			}
		})
	}
}
