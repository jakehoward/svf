package main

import (
	"errors"
 	"strings"
 	"strconv"
)

// Options is a collection of user definable program options
type Options struct {
	filepath string
	delimiter string
	writeFields []int
}

// OptionsBuilder builds up Options based on strings parsed from the command line
type OptionsBuilder struct {}

// Build takes strings representing fields and the delimiter and parses them into a
// domain specific Options struct, returning an error for invalid option values
func (b *OptionsBuilder) Build(delimiter string, fieldString string) (*Options, error) {
	options := new(Options)
	var err error
	if delimiter != "" && len(delimiter) == 1 {
		options.delimiter = delimiter
	} else {
		err = errors.New("Invalid delimiter, must be one character")
	}
	
	fields, fieldParseErr := parseWriteFields(fieldString); if fieldParseErr != nil {
		err = fieldParseErr
	} else {
		options.writeFields = fields
	}
	return options, err
}

func parseWriteFields(fieldString string) ([]int, error) {
	var err error
	if fieldString == "" {
		err = errors.New("Invalid list of fields, can't be empty")
	}
	fieldSymbols := strings.Split(fieldString, ",")
	var fields []int
	for _, symbol := range fieldSymbols {
		field, convErr := strconv.Atoi(symbol)
		if convErr != nil {
			err = errors.New("Invalid field list, contains non integer")
			break
		} else {
			// does this need the reassignment?
			fields = append(fields, field)
		}
	}
	return fields, err
}
