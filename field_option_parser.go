package main

import (
	"errors"
	"strconv"
	"strings"
)

// FieldOptionParser can read the raw field string of user cmd line options
// and translate it into the fields to be written on output
type FieldOptionParser struct{}

// Parse takes the raw string representing the fields to write out and
// turns it into either a list of fields or an error
func (f *FieldOptionParser) Parse(fieldString string) ([]int, error) {
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
			fields = append(fields, field)
		}
	}
	return fields, err
}
