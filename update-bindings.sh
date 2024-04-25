#!/bin/bash

if ! abigen -v &> /dev/null; then
    echo "abigen is not installed. Please run:"
    echo "go install github.com/ethereum/go-ethereum/cmd/abigen@latest"
    exit 1
fi

CURRENT_PATH=$(pwd)

cd /Users/f/Developer/nana/nana-suckers
git checkout master
git pull

echo "Updating BPSucker abi..."
forge inspect BPSucker abi > "$CURRENT_PATH/abi/BPSucker.json"
echo "Generating bindings..."
cd "$CURRENT_PATH"
abigen --abi="abi/BPSucker.json" --pkg=main --type BPSucker --out="BPSucker.go"

echo "Done!"
