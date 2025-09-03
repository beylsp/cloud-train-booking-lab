FROM mcr.microsoft.com/devcontainers/base:ubuntu

ARG PROTOC_VERSION=32.0
ARG BUF_VERSION=1.57.0

ARG BIN=/usr/local/bin

RUN apt-get update && \
    apt-get install -y \
        ca-certificates \
        unzip \
        curl \
        git \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN curl -SSLO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
    && unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local \
    && rm protoc-${PROTOC_VERSION}-linux-x86_64.zip

RUN curl -sSL https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-Linux-x86_64 -o ${BIN}/buf \
    && chmod +x ${BIN}/buf
