package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	// "runtime/pprof"

	flag "github.com/ogier/pflag"
)

/* Profiling code */
// var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
/* Profiling code */

var fieldString string
var fields []int
var delimiterString string

func init() {
	flag.StringVarP(&fieldString, "fields", "f", "", "fields to output, e.g. 1-3,7,9,14-20")
	flag.StringVarP(&delimiterString, "delimiter", "d", ",", "delimiter char on which to split fields")
	flag.Parse()
	var err error
	if fields, err = parseFieldString(fieldString); err != nil {
		log.Fatal(err)
	}
}

func main() {
	/* Profiling code */
	// if *cpuprofile != "" {
	// 	f, err := os.Create(*cpuprofile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	pprof.StartCPUProfile(f)
	// 	defer pprof.StopCPUProfile()
	// }
	/* Profiling code */

	i := inputSource(flag.Args())

	r := csv.NewReader(i)
	r.Comma = parseDelimiterString(delimiterString)
	r.FieldsPerRecord = -1 // Allow variable # of fields.

	w := csv.NewWriter(os.Stdout)
	w.Comma = r.Comma

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var outRecord []string
		if m := max(fields); m > len(record) {
			log.Fatalf("Can't output field %d, record only has length: %d, record: %v", m, len(record), record)
		}
		for _, i := range fields {
			outRecord = append(outRecord, record[i-1])
		}

		if err := w.Write(outRecord); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

func inputSource(args []string) io.Reader {
	var input io.Reader
	if len(args) == 0 {
		input = os.Stdin
	} else if len(args) == 1 {
		f, err := os.Open(args[0])
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		input = f
	} else {
		log.Fatalf("Please specify one argument for input file, %d specififed: %v", len(args), args)
	}
	return input
}

func parseDelimiterString(d string) rune {
	if len(d) != 1 {
		log.Fatalf("Delimiter can only be one character, got: %s", d)
	}
	return []rune(d)[0]
}

func max(xs []int) int {
	if len(xs) == 0 {
		log.Fatalf("Internal error: E1")
	}
	var max = xs[0]
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}
