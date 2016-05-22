package main

import (
	"testing"
)

var recordWriter = new(RecordWriter)
func TestRecordWriterCreatesSimpleRowCorrectly(t *testing.T) {
	record, _ := recordWriter.MakeRecord([]string{"1", "2", "3"}, ":")
	if record != "1:2:3" {
		t.Errorf("Expected record to be %s, but got %s", "1:2:3", record)
	}
}

func TestRecordWriterEscapesCorrectly(t *testing.T) {
	record, _ := recordWriter.MakeRecord([]string{"\"Hello\", said the man", "we've got a comma,", "no comma here"}, ",")
	expected := "\"\"\"Hello\"\", said the man\",\"we've got a comma,\",no comma here"
	if record != expected {
		t.Errorf("Expected record to be %s, but got %s", expected, record)
	}
}

func TestRoundRobinWithRecordParser(t *testing.T) {
	recordParser := new(RecordParser)
	start := "\"\"\"Hello\"\", said the man\",\"we've got a comma,\",no comma here"
	record, _ := recordParser.Parse(start, ",")
	end, _ := recordWriter.MakeRecord(record.fields, ",")
	if start != end {
		t.Errorf("Expected start and end to match but, start: %x, end: %x", start, end)
	}
}
	
