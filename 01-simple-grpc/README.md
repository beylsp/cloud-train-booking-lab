# Simple gRPC Client/Server

This chapter introduces the foundation of our Booking System by implementing a minimal **gRPC client and server**. The goal is to understand how services communicate using gRPC before adding more complexity in later chapters.

## Tools

All required tools are pre-installed in the **devcontainer**, but here’s what’s included and why:

- [protoc](https://protobuf.dev/) – the Protocol Buffers compiler
- Go plugins for protoc
    - protoc-gen-go – generates Go types from .proto files
    - protoc-gen-go-grpc – generates Go gRPC service stubs
- [buf.build](https://buf.build/) - used to manage .proto files, lint them, and generate code consistently

You don’t need to install these locally — just open the chapter in the devcontainer and you’re ready to go.
