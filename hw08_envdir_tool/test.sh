#!/usr/bin/env bash
set -xeuo pipefail

go build -o go-envdir

export HELLO="SHOULD_REPLACE"
export FOO="SHOULD_REPLACE"
export UNSET="SHOULD_REMOVE"

result=$(./go-envdir "$(pwd)/testdata/env" "/bin/bash" "$(pwd)/testdata/echo.sh" arg1=1 arg2=2)
expected='HELLO is ("hello")
BAR is (bar)
FOO is (   foo
with new line)
UNSET is ()
arguments are arg1=1 arg2=2'

[ "${result}" = "${expected}" ] || (echo -e "invalid output: ${result}" && exit 1)

rm -f go-envdir
echo "PASS"
