#!/usr/bin/env bash

_assertEqual () {
    expected="$1"
    actual="$2"
    if [ "$expected" != "$actual" ]; then
	echo "Expected: $expected, but actually got: $actual" 
	exit 1
    fi
}

input="Hello,fine,world\nHow,are,you?"

# Test svf acts as identity function when given no options
output="$(echo ${input} | svf)"
_assertEqual "$input" "$output"

# Test svf reads file when passed as argument
test_file_extension="$(date +%s)"
test_input_file="test_input.${test_file_extension}"
echo "$input" > "$test_input_file"
output2="$(svf ${test_input_file})"
rm "$test_input_file"
_assertEqual "$input" "$output2"

