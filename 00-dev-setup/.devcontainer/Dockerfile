FROM mcr.microsoft.com/devcontainers/base:ubuntu

RUN apt-get update && \
    apt-get install -y \
        ca-certificates \
        curl \
        git && \
    rm -rf /var/lib/apt/lists/*
