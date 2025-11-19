#!/bin/bash

set -e

BINARY_NAME="taskgo"
INSTALL_DIR="/usr/local/bin"

echo "Installing $BINARY_NAME..."

# Build the project
# Check for Go
if command -v go &> /dev/null; then
    echo "Go is already installed."
else
    echo "Go not found. Installing..."
    GO_VERSION=$(curl -s https://go.dev/VERSION?m=text | head -n 1)
    echo "Downloading Go version: $GO_VERSION"
    wget "https://go.dev/dl/${GO_VERSION}.linux-amd64.tar.gz"
    sudo tar -C /usr/local -xzf "${GO_VERSION}.linux-amd64.tar.gz"
    rm "${GO_VERSION}.linux-amd64.tar.gz"
    export PATH=$PATH:/usr/local/go/bin
fi

echo "Building binary..."

# Check if we are in the project root
if [ ! -f "go.mod" ]; then
    echo "go.mod not found. Cloning repository..."
    TEMP_DIR=$(mktemp -d)
    git clone https://github.com/MohakGupta2004/taskgo.git "$TEMP_DIR"
    cd "$TEMP_DIR"
fi

go build -o $BINARY_NAME main.go

echo "Moving binary to $INSTALL_DIR..."
sudo mv $BINARY_NAME $INSTALL_DIR/

echo "Installation complete! You can now run '$BINARY_NAME'."
