#!/bin/sh

cd goapi
echo "run goimports"
goimports -w .

echo "run gofmt"
gofmt -w .

echo "run golint"
golint ./...

echo "run gsc"
gsc ./...

echo "run gosec"
gosec ./...

echo "run staticcheck"
staticcheck ./gameapi/...

echo "run errcheck"
errcheck ./...

echo "run misspell"
misspell .

cd ..