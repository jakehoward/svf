#!/usr/bin/env bash

set -e

_assertEqual () {
    expected="$1"
    actual="$2"
    name="$3"
    if [ "$expected" != "$actual" ]; then
	printf "FAIL: %s\n", "${name}"
	printf "Expected: '%s' but actually got: '%s'\n" "$expected" "$actual"
	exit 1
    fi
}

# Basic Functionality
testname="Basic Functionality"
input="1,2,3"
output="$(echo ${input} | svf -d ',' -f 1,3)"
_assertEqual "1,3" "$output" "$testname"

# Different Delimiter
testname="Semi Colon Delimiter"
input=$'test:a new delimiter:passed:hopefully'
output="$(echo ${input} | svf -d ':' -f 1,3,4)"
_assertEqual $'test:passed:hopefully' "$output" "$testname"

# Newlines
cat functional_test_input.csv | svf -d ',' -f 1,2,3,4 > out1.tmp
svf -d ',' -f 1,2,3,4 out1.tmp > out2.tmp
diff out1.tmp out2.tmp
diff out1.tmp functional_test_expected_output.csv

rm out1.tmp out2.tmp

echo "Functional Tests: PASS"
exit 0
