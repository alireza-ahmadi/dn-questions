package dnquestions

import (
	"testing"
)

func TestRepeatedString(t *testing.T) {
	tests := []struct {
		str      string
		num      int
		expected int
	}{
		{"abcac", 10, 4},
		{"foobar", 100, 16},
		{"bruxa", 4, 0},
		{"something", 1000, 0},
		{"", 100, 0},
		{"a", 1000000, 1000000},
	}

	for _, testCase := range tests {
		output := RepeatedString(testCase.str, testCase.num)
		if output != testCase.expected {
			t.Errorf("Case %s, Expected %d, Returned %d", testCase.str, testCase.expected, output)
		}
	}
}

func TestMinimumRemovals(t *testing.T) {
	tests := []struct {
		str      string
		expected int
	}{
		{"AAAA", 3},
		{"BBBBB", 4},
		{"ABABABAB", 0},
		{"BABABA", 0},
		{"AAABBB", 4},
		{"A", 0},
		{"B", 0},
		{"AB", 0},
		{"BAB", 0},
		{"ABB", 1},
		{"", 0},
	}

	for _, testCase := range tests {
		output := MinimumRemovals(testCase.str)
		if output != testCase.expected {
			t.Errorf("Case %s, Expected %d, Returned %d", testCase.str, testCase.expected, output)
		}
	}
}

func TestMinSwapsToSort(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{1, 5, 4, 3, 2}, 2},
		{[]int{4, 3, 1, 2}, 3},
		{[]int{2, 3, 4, 1, 5}, 3},
		{[]int{1, 3, 5, 2, 4, 6, 7}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 0},
		{[]int{7, 1, 3, 2, 4, 5, 6}, 5},
		{[]int{7, 6, 5, 4, 3, 2, 1}, 3},
		{[]int{1}, 0},
		{[]int{}, 0},
	}

	for idx, testCase := range tests {
		output := MinSwapsToSort(testCase.arr)
		if output != testCase.expected {
			t.Errorf("Case %d, Expected %d, Returned %d", idx, testCase.expected, output)
		}
	}
}
