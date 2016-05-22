package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	flag "github.com/ogier/pflag"
)

func main() {
	optionsBuilder := new(OptionsBuilder)
	run(optionsBuilder)
}

func run(optionsBuilder *OptionsBuilder) {
 	var delimiter = flag.StringP("delimiter", "d", ",", "Character used to split fields")
	var fields = flag.StringP("fields", "f", "", "Fields you would like printed to stdout")
	flag.Parse()
	
	options, err := optionsBuilder.Build(*delimiter, *fields)
	if err != nil {
		os.Exit(1)
	}
	
	input := os.Stdin
	if options.filepath != "" {
		file, err := os.Open(options.filepath)
		if err != nil {
			os.Exit(1)
		}
		input = file
	}
	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(os.Stdout)
	recordParser := new(RecordParser)
	for scanner.Scan() {
		var fieldsToWrite []string
		record, err := recordParser.Parse(scanner.Text(), options.delimiter); if err != nil {
			fmt.Println("Error occurred parsing record")
			os.Exit(1);
		}
		for index, field := range record.fields {
			if contains(options.writeFields, index + 1) {
				fieldsToWrite = append(fieldsToWrite, field)
			}
		}
		fmt.Fprintln(writer, strings.Join(fieldsToWrite, options.delimiter))
		writer.Flush()
	}
}

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
