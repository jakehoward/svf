package main

import (
	"reflect"
	"testing"
)

func TestFieldStringParser(t *testing.T) {
	tests := []struct {
		fieldString string
		expected    []int
	}{
		{"", []int{}},
		{"1", []int{1}},
		{"3,4", []int{3, 4}},
		{"1,457,2", []int{1, 457, 2}},
	}
	var actual []int
	for _, test := range tests {
		actual, _ = parseFieldString(test.fieldString)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Expected: %v, Actual: %v", test.expected, actual)
		}
		actual = nil
	}
}
