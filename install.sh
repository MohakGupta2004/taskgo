#!/bin/bash

set -e

BINARY_NAME="taskgo"
INSTALL_DIR="/usr/local/bin"

echo "Installing $BINARY_NAME..."

# Build the project
if ! command -v go &> /dev/null; then
    GO_VERSION=$(curl -s https://go.dev/VERSION?m=text | head -n 1)
    echo "Downloading Go version: $GO_VERSION"
    wget "https://go.dev/dl/${GO_VERSION}.linux-amd64.tar.gz"
    sudo tar -C /usr/local -xzf "${GO_VERSION}.linux-amd64.tar.gz"
    rm "${GO_VERSION}.linux-amd64.tar.gz"
    export PATH=$PATH:/usr/local/go/bin
fi

echo "Building binary..."
go build -o $BINARY_NAME main.go

echo "Moving binary to $INSTALL_DIR..."
mv $BINARY_NAME $INSTALL_DIR/

echo "Installation complete! You can now run '$BINARY_NAME'."
