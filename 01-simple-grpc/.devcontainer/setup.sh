#!/bin/bash
set -euo pipefail

echo "🔄 Installing Go plugins for protoc"

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

CLUSTER_NAME="${CLUSTER_NAME:-dev-cluster}"

echo "🔄 Resetting kind cluster: $CLUSTER_NAME"

# Delete if exists
kind delete cluster --name "$CLUSTER_NAME" || true

# Create fresh cluster
kind create cluster --name "$CLUSTER_NAME"
