#!/bin/zsh

sed -i.bak '/- deadcode/d' .golangci.yml
sed -i '' '/- unused/d' .golangci.yml
sed -i '' '/- structcheck/d' .golangci.yml

for d in $(ls)
do
  if [[ $d == hw* ]]; then
    cd $d
    echo "Lint ${d}..."
    golangci-lint run ./...
    cd ..
  fi
done

mv .golangci.yml.bak .golangci.yml
