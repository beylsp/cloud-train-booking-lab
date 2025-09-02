#!/bin/bash
set -euo pipefail

CLUSTER_NAME="${CLUSTER_NAME:-dev-cluster}"

echo "🔄 Resetting kind cluster: $CLUSTER_NAME"

# Delete if exists
kind delete cluster --name "$CLUSTER_NAME" || true

# Create fresh cluster
kind create cluster --name "$CLUSTER_NAME"
