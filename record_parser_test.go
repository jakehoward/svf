package main

import (
	"testing"
	"reflect"
)


var recParser = new(RecordParser)

func TestRecordParserReturnsCorrectRecordGivenSimpleInput(t *testing.T) {
	record, _ := recParser.Parse("this,is,a,1,2,simple,test", ",")
	expected := []string{"this", "is", "a", "1", "2", "simple", "test"}
	got := record.fields
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestRecordParserReturnsCorrectRecordGivenEscapedDelimiter(t *testing.T) {
	record, _ := recParser.Parse("1,2,\"escape, from you\"", ",")
	expected := []string{"1", "2", "escape, from you"}
	got := record.fields
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestRecordParserReturnsCorrectRecordGivenEscapedEscapeChars(t *testing.T) {
	record, _ := recParser.Parse(",hi,\"\"\"Hello\"\" said the man\",bye", ",")
	expected := []string{"","hi","\"Hello\" said the man", "bye"}
	got := record.fields
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestRecordParserReturnsCorrectRecordGivenEscapedEscapeAndDelimiterChars(t *testing.T) {
	record, _ := recParser.Parse(",hi,\"\"\"Goodbye\"\", said the other chap\",cheerio", ",")
	expected := []string{"","hi","\"Goodbye\", said the other chap", "cheerio"}
	got := record.fields
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestEdgeEscapeCaseEmptyQuotes(t *testing.T) {
	record, _ := recParser.Parse("a,\"\"\"\"\"\",b", ",")
	expected := []string{"a", "\"\"", "b"}
	got := record.fields
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestEdgeEscapeCaseEmptyEscapedField(t *testing.T) {
	record, _ := recParser.Parse("a,\"\",b", ",")
	expected := []string{"a", "", "b"}
	got := record.fields
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestEdgeCaseEscapeEscapedQuotesOnlyField(t *testing.T) {
	record, _ := recParser.Parse("\"\"\"\"\"\"", ",")
	expected := []string{"\"\""}
	got := record.fields
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
func TestEdgeCaseEscapeEscapedQuotesLastField(t *testing.T) {
	record, _ := recParser.Parse("a,\"\"\"\"\"\"", ",")
	expected := []string{"a", "\"\""}
	got := record.fields
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
func TestEdgeCaseEscapeEscapedQuotesFirstField(t *testing.T) {
	record, _ := recParser.Parse("\"\"\"\"\"\",b", ",")
	expected := []string{"\"\"", "b"}
	got := record.fields
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestEdgeCaseEscapedQuotesGalore(t *testing.T) {
	// Record: """Quotes, ""are"" fun"""
	record, _ := recParser.Parse("\"\"\"Quotes, \"\"are\"\" fun\"\"\"", ",")
	expected := []string{"\"Quotes, \"are\" fun\""}
	got := record.fields
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
