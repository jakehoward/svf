package main

import (
	"testing"
)

func TestArgumentParserReadsFilepathWhenOnlyArgument(t *testing.T) {
	filepath := "../data/some_file.csv"
	rawCommandLineOptions := []string{filepath}
	clp := new(CommandLineOptionsParser)
	opts := clp.Parse(rawCommandLineOptions)
	if opts.filepath != filepath {
		t.Errorf("Expected filepath to be: %q, but actually got: %q",filepath, opts.filepath)
	}
}

func TestArgumentParserDoesNotBlowUpWhenGivenZeroOptions(t *testing.T) {
	clp := new(CommandLineOptionsParser)
	clp.Parse([]string{})
}
