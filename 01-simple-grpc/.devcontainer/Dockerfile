FROM mcr.microsoft.com/devcontainers/base:ubuntu

RUN apt-get update && \
    apt-get install -y \
        ca-certificates \
        curl \
        git && \
    rm -rf /var/lib/apt/lists/*

# Install buf CLI
ARG BUF_VERSION=1.57.0
RUN curl -sSL "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-Linux-x86_64" -o /usr/local/bin/buf \
    && chmod +x /usr/local/bin/buf
