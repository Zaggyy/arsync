#!/bin/sh

GOOS=linux GOARCH=arm64 go build

echo "Built for Raspberry Pi"
