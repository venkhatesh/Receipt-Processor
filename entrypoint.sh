#!/bin/sh

# To Run Test cases
echo "Running tests..."
go test -v ./...
if [ $? -ne 0 ]; then
    echo "Tests failed. Exiting..."
    exit 1
fi

# To Start GO Services
echo "Tests passed. Starting the service..."
export VERBOSE=true
exec go run .
