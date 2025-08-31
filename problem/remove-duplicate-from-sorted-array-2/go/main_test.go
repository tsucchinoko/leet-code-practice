package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		expectedK    int
		expectedNums []int
	}{
		{
			name:         "Example 1",
			input:        []int{1, 1, 1, 2, 2, 3},
			expectedK:    5,
			expectedNums: []int{1, 1, 2, 2, 3},
		},
		{
			name:         "Example 2",
			input:        []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expectedK:    7,
			expectedNums: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			name:         "Empty array",
			input:        []int{},
			expectedK:    0,
			expectedNums: []int{},
		},
		{
			name:         "Single element",
			input:        []int{1},
			expectedK:    1,
			expectedNums: []int{1},
		},
		{
			name:         "Two identical elements",
			input:        []int{1, 1},
			expectedK:    2,
			expectedNums: []int{1, 1},
		},
		{
			name:         "Two different elements",
			input:        []int{1, 2},
			expectedK:    2,
			expectedNums: []int{1, 2},
		},
		{
			name:         "All unique elements",
			input:        []int{1, 2, 3, 4, 5},
			expectedK:    5,
			expectedNums: []int{1, 2, 3, 4, 5},
		},
		{
			name:         "All elements are same, more than 2",
			input:        []int{5, 5, 5, 5, 5},
			expectedK:    2,
			expectedNums: []int{5, 5},
		},
		{
			name:         "Longer example with various duplicates",
			input:        []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5},
			expectedK:    10,
			expectedNums: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		},
		{
			name:         "Leading and trailing duplicates",
			input:        []int{1, 1, 1, 2, 3, 3, 3},
			expectedK:    5,
			expectedNums: []int{1, 1, 2, 3, 3},
		},
		{
			name:         "Negative numbers",
			input:        []int{-2, -2, -1, 0, 0, 0, 1, 1},
			expectedK:    7,                             // ここを 6 から 7 に修正
			expectedNums: []int{-2, -2, -1, 0, 0, 1, 1}, // ここを [-2, -2, -1, 0, 0, 1] から修正
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// テストの独立性の担保のため、inputCopyにtt.inputをコピーする
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			gotK := removeDuplicates(inputCopy)

			if gotK != tt.expectedK {
				t.Errorf("removeDuplicates() for %s returned k = %d, want %d", tt.name, gotK, tt.expectedK)
			}

			if !reflect.DeepEqual(inputCopy[:gotK], tt.expectedNums) {
				t.Errorf("removeDuplicates() for %s modified nums to %v, want %v", tt.name, inputCopy[:gotK], tt.expectedNums)
			}
		})
	}
}
