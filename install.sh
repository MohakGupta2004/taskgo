#!/bin/bash

set -e

BINARY_NAME="taskgo"
INSTALL_DIR="/usr/local/bin"

echo "Installing $BINARY_NAME..."

# Build the project
if ! command -v go &> /dev/null; then
    wget "https://dl.google.com/go/$(curl https://go.dev/VERSION?m=text).linux-amd64.tar.gz"
    tar -C /usr/local -xzf "$(curl https://go.dev/VERSION?m=text).linux-amd64.tar.gz"
    export PATH=$PATH:/usr/local/go/bin
fi

echo "Building binary..."
go build -o $BINARY_NAME main.go

echo "Moving binary to $INSTALL_DIR..."
mv $BINARY_NAME $INSTALL_DIR/

echo "Installation complete! You can now run '$BINARY_NAME'."
