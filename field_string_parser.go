package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseFieldString(fs string) ([]int, error) {
	var fieldNums = []int{}
	if len(fs) == 0 {
		return fieldNums, nil
	}
	fields := strings.Split(fs, ",")
	for _, f := range fields {
		i, err := strconv.Atoi(f)
		if err != nil {
			return nil, fmt.Errorf("Fields can only be numerical values, got: %s", f)
		}
		fieldNums = append(fieldNums, i)
	}
	return fieldNums, nil
}
