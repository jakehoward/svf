#!/usr/bin/env bash

_assertEqual () {
    expected="$1"
    actual="$2"
    if [ "$expected" != "$actual" ]; then
	echo "FAIL"
	echo "Expected: $expected, but actually got: $actual" 
	exit 1
    fi
}

# Test svf lets user choose delimeter and fields to print
colon_input="1:2:3"
output="$(echo ${colon_input} | svf -d ':' -f 1,3)"
_assertEqual "1:3" "$output"

