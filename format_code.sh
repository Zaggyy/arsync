#!/bin/sh
# Check if we are in the right directory
if [ ! -d "client" ] || [ ! -d "server" ]; then
    echo "Please run this script from the root of the project"
    exit 1
fi

# Format the code
echo "Formatting code..."

cd client || exit
gofmt -w ./*.go

cd ../server || exit
gofmt -w ./*.go

cd .. || exit

echo "Done"
