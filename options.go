package main

// Options is a collection of user definable program options
type Options struct {
	filepath string
}

// OptionsParser represents the ability to parse program options
type OptionsParser interface {
	Parse([]string) *Options
}

// CommandLineOptionsParser is an OptionsParser designed to parse options sent as command line flags
type CommandLineOptionsParser struct {}

// Parse takes command line options and returns a domain specific object representing thier meaning
func (c CommandLineOptionsParser) Parse(rawOptions []string) *Options {
	options := new(Options)
	if len(rawOptions) > 0 {
		options.filepath = rawOptions[0]
	}
	return options
}

