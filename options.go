package main

import (
	"errors"
	"io"
	"os"
)

// Options is a collection of user definable program options
type Options struct {
	inputSource io.Reader
	delimiter string
	writeFields []int
}

// OptionsBuilder builds up Options based on strings parsed from the command line
type OptionsBuilder struct {}

// Build takes strings representing fields and the delimiter and parses them into a
// domain specific Options struct, returning an error for invalid option values
func (b *OptionsBuilder) Build(delimiter string, fieldString string, filepath string) (*Options, error) {
	options := new(Options)
	var err error
	if delimiter != "" && len(delimiter) == 1 {
		options.delimiter = delimiter
	} else {
		err = errors.New("Invalid delimiter, must be one character")
	}
	fieldParser := new(FieldOptionParser)
	fields, fieldParseErr := fieldParser.Parse(fieldString); if fieldParseErr != nil {
		err = fieldParseErr
	} else {
		options.writeFields = fields
	}

	input := os.Stdin
	if filepath != "" {
		file, fileErr := os.Open(filepath)
		if fileErr != nil {
			err = errors.New("failed to open file at: " + filepath)
		}
		input = file
	}
	options.inputSource = input
	return options, err
}
