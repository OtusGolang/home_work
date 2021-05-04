#!/bin/bash

dst=$1
if [[ ! -d "${dst}" ]]; then
    echo "Usage: ./.scripts/sync.sh <destination dir>."
    echo "The destination dir should exist"
    exit 1
fi

GLOBIGNORE=".:..:.git"
for f in *; do
    [[ -d "${dst}/${f}" ]] && [[ ! -f "${dst}/${f}/.sync" ]] && continue

    echo "syncing ${f}..."
    cp -R "${f}" "${dst}"
done
