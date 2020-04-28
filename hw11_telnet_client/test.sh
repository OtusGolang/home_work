#!/usr/bin/env bash
set -xeuo pipefail

go build -o go-telnet

(echo -e "Hello\nFrom\nNC\n" && cat 2>/dev/null) | nc -l localhost 4242 >/tmp/nc.out &
NC_PID=$!

sleep 1
echo -e "I\nam\nTELNET client\n" | ./go-telnet --timeout=5s localhost 4242 >/tmp/telnet.out 2>/tmp/telnet.err

kill ${NC_PID} 2>/dev/null || true

function fileEquals() {
  local fileData
  fileData=$(cat "$1")
  [ "${fileData}" = "${2}" ] || (echo -e "unexpected output, $1:\n${fileData}" && exit 1)
}

expected_nc_out='I
am
TELNET client'
fileEquals /tmp/nc.out "${expected_nc_out}"

expected_telnet_out='Hello
From
NC'
fileEquals /tmp/telnet.out "${expected_telnet_out}"

expected_telnet_err='...Connected to localhost:4242
...EOF'
fileEquals /tmp/telnet.err "${expected_telnet_err}"

rm -f go-telnet
echo "PASS"
