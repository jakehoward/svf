package main

import (
	"errors"
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
		is, err := parseFieldToken(f)
		if err != nil {
			return nil, err
		}
		for _, i := range is {
			if i == 0 {
				return nil, errors.New("0 is not a valid field index, fields is 1 based")
			}
			fieldNums = append(fieldNums, i)
		}
	}
	return fieldNums, nil
}

func parseFieldToken(s string) ([]int, error) {
	var nums = []int{}
	if len(s) == 0 {
		return nums, nil
	}

	if !strings.Contains(s, "-") {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("Unable to parse '%s' as number in field string", s)
		}
		return []int{num}, nil
	}

	se := strings.Split(s, "-")
	var err error
	if len(se) == 2 {
		if se[1] == "" {
			return nil, fmt.Errorf("Field string syntax n- not yet supported. Offending section: %s-", se[0])
		}

		start, err := strconv.Atoi(se[0])
		if err != nil {
			return nil, fmt.Errorf("Unable to parse '%s' as number in field string", se[0])
		}

		end, err := strconv.Atoi(se[1])
		if err != nil {
			return nil, fmt.Errorf("Unable to parse '%s' as number in field string", se[1])
		}
		if start == end || start > end {
			return nil, fmt.Errorf("Start must be smaller than end in field string range, start: %d, end: %d", start, end)
		}
		nums = makeRangeIncl(start, end)
	}
	return nums, err
}

func makeRangeIncl(s, e int) []int {
	var nums = []int{}
	for i := s; i <= e; i++ {
		nums = append(nums, i)
	}
	return nums
}
