#!/bin/sh

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

echo "Built for Windows"
