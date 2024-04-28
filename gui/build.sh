#!/bin/sh

CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build

echo "Built for Windows"
