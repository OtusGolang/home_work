#!/usr/bin/env bash
set -xeuo pipefail

expected='!SUTO ,olleH'
result=$(go run main.go | sed 's/^ *//;s/ *$//')
[ "${result}" = "${expected}" ] || (echo -e "invalid output: ${result}" && exit 1)

echo "PASS"
