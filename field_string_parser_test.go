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
		{"1-3", []int{1, 2, 3}},
	}
	var actual []int
	var err error
	for _, test := range tests {
		if actual, err = parseFieldString(test.fieldString); err != nil {
			t.Errorf("Expected: %v, but got error %v", test.expected, err)
			break
		}
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Expected: %v, Actual: %v", test.expected, actual)
		}
		actual, err = nil, nil
	}
}

func TestFieldStringRangeParser(t *testing.T) {
	tests := []struct {
		rangeString string
		expected    []int
	}{
		{"", []int{}},
		{"1", []int{1}},
		{"1-3", []int{1, 2, 3}},
		{"1-2", []int{1, 2}},
		{"120-125", []int{120, 121, 122, 123, 124, 125}},
	}
	var actual []int
	var err error
	for _, test := range tests {
		if actual, err = parseFieldToken(test.rangeString); err != nil {
			t.Errorf("Expected %v, but got error %v", test.expected, err)
			break
		}
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Expected: %v, Actual: %v", test.expected, actual)
		}
		actual, err = nil, nil
	}
}

func TestFieldStringParserErrors(t *testing.T) {
	tests := []struct {
		fieldString string
	}{
		{"3-2"},
		{"1-"},
		{"5-a"},
		{"z-5"},
		{"1,2 ,4"},
		{"1,0,3-4"},
	}
	var err error
	for _, test := range tests {
		if _, err = parseFieldString(test.fieldString); err == nil {
			t.Errorf("Expected error but error was nil for input: %v", test.fieldString)
		}
		err = nil
	}
}
