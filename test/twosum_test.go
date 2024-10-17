package main

import (
	"reflect"
	"testing"

	"github.com/reaper2171/find-pairs/utils"
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
			expected: [][]int{{0, 3}, {1, 3}},
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
			expected: [][]int{{1, 2}, {0, 3}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := utils.TwoSum(tc.numbers, tc.target)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("TwoSum(%v, %d) = %v; want %v", tc.numbers, tc.target, result, tc.expected)
			}
		})
	}
}
