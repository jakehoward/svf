#!/usr/bin/env bash

set -e

_assertEqual () {
    expected="$1"
    actual="$2"
    if [ "$expected" != "$actual" ]; then
	echo "FAIL"
	echo "Expected: $expected, but actually got: $actual" 
	exit 1
    fi
}

# Example test 
# colon_input="1:2:3"
# output="$(echo ${colon_input} | svf -d ':' -f 1,3)"
# _assertEqual "1:3" "$output"

echo "PASS"
exit 0
