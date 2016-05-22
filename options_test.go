package main

import (
	"testing"
	"reflect"
	"os"
)
/*
 * OptBuilder currently returns the options even if there's an error, should probably change that
* Opt builder probably needs to delegate field parsing to a standalone type/testable entity
 */
var optBuilder = new(OptionsBuilder)

func TestOptionsBuilderReturnsOptionsWithDelimiter(t *testing.T) {
	delim := ":"
	options, _ := optBuilder.Build(delim, "", "")
	if options.delimiter != delim {
		t.Errorf("Expected options to have delimiter: %q, but actually had: %q", delim, options.delimiter)
	}
}

func TestOptionsBuilderReturnsErrorIfPresentedWithEmptyDelimiter(t *testing.T) {
	delim := ""
	_, err := optBuilder.Build(delim, "", "")
	if err == nil {
		t.Errorf("Expected error when passing empty delimiter option builder")
	}
}

func TestOptionsBuilderReturnsErrorIfPassedADelimiterLongerThanOneCharacter(t *testing.T) {
	delim := "::"
	_, err := optBuilder.Build(delim, "", "")
	if err == nil {
		t.Errorf("Expected error when passing multi-char delimiter to option builder")
	}
}

func TestOptionsBuilderReturnsFields(t *testing.T) {
	fieldString := "3,6"
	options, _ := optBuilder.Build(",", fieldString, "")
	if !reflect.DeepEqual(options.writeFields, []int{3, 6}) {
		t.Errorf("Expected write fields to be %x, but they were actually %x", []int{3,6}, options.writeFields)
	}
}

func TestOptionsBuilderReturnsStdinAsInputIfNoFilepathProvided(t *testing.T) {
	filepath := ""
	options, _ := optBuilder.Build(",", "1", filepath)
	if !reflect.DeepEqual(options.inputSource, os.Stdin) {
		t.Errorf("Expected input source to be stdin")
	}
}

func TestOptionsBuilderReturnsErrWhenFilepathPointsNowhere(t *testing.T) {
	filepath := "thisfiledoesnotexist.txt"
	_, err := optBuilder.Build(",", "1", filepath)
	if err == nil {
		t.Errorf("Expected error when passing non existent filepath")
	}
}

// Figure out how to test file handle returned
