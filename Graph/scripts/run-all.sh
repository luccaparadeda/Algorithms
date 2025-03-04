#!/bin/bash

BASE_DIR="$(pwd)"

find "$BASE_DIR" -type f -name "main.go" | while read -r file; do
    dir=$(dirname "$file")
    echo "Running in: $dir"

    go run $dir/main.go
done
