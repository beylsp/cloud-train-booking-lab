#!/bin/bash

set -euo pipefail

if [ $# -ne 2 ]; then
  echo "Usage: $0 <old-dir> <new-dir>"
  exit 1
fi

OLD_DIR=${1%/}
NEW_DIR=${2%/}

if [ ! -d "$OLD_DIR" ]; then
  echo "Error: directory '$OLD_DIR' does not exist."
  exit 1
fi

if [ -e "$NEW_DIR" ]; then
  echo "Error: target '$NEW_DIR' already exists."
  exit 1
fi

rsync -av \
    --exclude 'gen' \
    --exclude 'build' \
    "$OLD_DIR/" "$NEW_DIR/"

find "$NEW_DIR" -type f -exec sed -i "s/$OLD_DIR/$NEW_DIR/g" {} + 

echo "Created '$NEW_DIR' from '$OLD_DIR'."
