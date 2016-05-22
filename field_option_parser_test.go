package main

import (
	"testing"
	"reflect"
)

var fieldParser = new(FieldOptionParser)
func TestFieldParserReturnsCorrectFieldWhenPassedOneField(t *testing.T) {
	fieldString := "1"
	fields, _ := fieldParser.Parse(fieldString)
	if !reflect.DeepEqual(fields, []int{1}) {
		t.Errorf("Expected options to have write fields %x, but actually got %x", []int{1}, fields)
	}
}

func TestFieldParserReturnsErrorWhenPassedEmtpyFieldList(t *testing.T) {
	fieldString := ""
	_, err := fieldParser.Parse(fieldString)
	if err == nil {
		t.Errorf("Expected option builder to return an error when passed an empty field string")
	}
}

func TestFieldParserReturnsErrorWhenPassedInvalidContent(t *testing.T) {
	fieldString := "a"
	_, err := fieldParser.Parse(fieldString)
	if err == nil {
		t.Errorf("Expected option builder to return an error when passed an invalid field string")
	}
}

func TestFieldParserHandlesMultipleFields(t *testing.T) {
	fieldString := "3,6"
	fields, _ := fieldParser.Parse(fieldString)
	if !reflect.DeepEqual(fields, []int{3, 6}) {
		t.Errorf("Expected write fields to be %x, but they were actually %x", []int{3,6}, fields)
	}
}
