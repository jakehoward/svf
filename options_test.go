package main

import (
	"testing"
	"reflect"
)
/*
 * OptBuilder currently returns the options even if there's an error, should probably change that
* Opt builder probably needs to delegate field parsing to a standalone type/testable entity
 */
func TestOptionsBuilderReturnsOptionsWithDelimiter(t *testing.T) {
	optBuilder := new(OptionsBuilder)
	delim := ":"
	options, _ := optBuilder.Build(delim, "")
	if options.delimiter != delim {
		t.Errorf("Expected options to have delimiter: %q, but actually had: %q", delim, options.delimiter)
	}
}

func TestOptionsBuilderReturnsErrorIfPresentedWithEmptyDelimiter(t *testing.T) {
	optBuilder := new(OptionsBuilder)
	delim := ""
	_, err := optBuilder.Build(delim, "")
	if err == nil {
		t.Errorf("Expected error when passing empty delimiter option builder")
	}
}

func TestOptionsBuilderReturnsErrorIfPassedADelimiterLongerThanOneCharacter(t *testing.T) {
	optBuilder := new(OptionsBuilder)
	delim := "::"
	_, err := optBuilder.Build(delim, "")
	if err == nil {
		t.Errorf("Expected error when passing multi-char delimiter to option builder")
	}
}

func TestOptionsBuilderReturnsFields(t *testing.T) {
	optBuilder := new(OptionsBuilder)
	fieldString := "3,6"
	options, _ := optBuilder.Build(",", fieldString)
	if !reflect.DeepEqual(options.writeFields, []int{3, 6}) {
		t.Errorf("Expected write fields to be %x, but they were actually %x", []int{3,6}, options.writeFields)
	}
}
