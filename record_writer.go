package main

import (
	"strings"
)

// RecordWriter makes a record given fields
type RecordWriter struct {}

// MakeRecord takes a list of fields and a delimiter and returns a string joining them and escaping special characters
func (r *RecordWriter) MakeRecord(fields []string, delimiter string) (string, error) {
	var err error
	escapedFields := []string{}
	for _, field := range fields {
		escapedFields = append(escapedFields, escapeField(field, delimiter))
	}
	return strings.Join(escapedFields, delimiter), err	
}

func escapeField(field string, delimiter string) string {
	var escapedField = strings.Replace(field, "\"", "\"\"", -1)
	if strings.Contains(field, delimiter) || strings.Contains(field, "\"") {
		escapedField = "\"" + escapedField + "\""
	}
	return escapedField
}
