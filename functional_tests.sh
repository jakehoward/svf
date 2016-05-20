#!/usr/bin/env bash

_assertEqual () {
    expected="$1"
    actual="$2"
    if [ "$expected" != "$actual" ]; then
	echo "Expected: $expected, but actually got: $actual" 
	exit 1
    fi
}

# Test svf acts as identity function when given no options
input="Hello,fine,world\nHow,are,you?"
output="$(echo ${input} | svf)"
_assertEqual "$input" "$output"

