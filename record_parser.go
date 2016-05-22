package main

import (
	"errors"
)

// Record represents a row in a *sv file
type Record struct {
	fields []string
}

// RecordParser takes the raw representation of a *sv record and parses it into a list of fields
type RecordParser struct {}

// Parse takes a row of *sv file and, taking into account escaping rules, returns a record representing the row
func (r *RecordParser) Parse(recordString string, delimiter string) (*Record, error) {
	var err error
	record := new(Record)

	var fields []string
	var field = []rune{}
	escapeRune := '"'
	delimiterRune := []rune(delimiter)[0]
	var inEscape = false
	recordStringRunes := stringToRuneSlice(recordString)
	for i, c := range recordStringRunes {
		if c == escapeRune {
			oneAhead, err := peekAheadOne(recordStringRunes, i); if err == nil {
				if inEscape && oneAhead == escapeRune {
					field = append(field, c)
				}
			}
			inEscape = !inEscape
		} else if c == delimiterRune && !inEscape {
			fields = append(fields, string(field))
			field = []rune{}
		} else {
			field = append(field, c)
		}
	}
	fields = append(fields, string(field))
	
	record.fields = fields
	return record, err
}

func peekAheadOne(rs []rune, index int) (rune, error) {
	var err error
	var value rune 
	if len(rs) > index + 1 {
		value = rs[index + 1]
	} else {
		err = errors.New("Can't peek ahead, at end of rune slice")
	}
	return value, err
}

func stringToRuneSlice(s string) []rune {
	r := []rune{}
	for _, rune := range s {
		r = append(r, rune)
	}
	return r
}
