package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	optionsParser := new(CommandLineOptionsParser)
	run(os.Args[1:], optionsParser)
}

func run(rawOptions []string, optionsParser OptionsParser) {
	options := optionsParser.Parse(rawOptions)
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
	for scanner.Scan() {
		fmt.Fprintln(writer, scanner.Text())
		writer.Flush()
	}
}
