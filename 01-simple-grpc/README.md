# 01-simple-grpc

This chapter introduces the foundation of our Booking System by implementing a minimal **gRPC client and server**. The goal is to understand how services communicate using gRPC before adding more complexity in later chapters.

## Tools

All required tools are pre-installed in the **devcontainer**, but here’s what’s included and why:

- [protoc](https://protobuf.dev/) – the Protocol Buffers compiler
- Go plugins for protoc
    - protoc-gen-go – generates Go types from .proto files
    - protoc-gen-go-grpc – generates Go gRPC service stubs
- [buf.build](https://buf.build/) - used to manage .proto files, lint them, and generate code consistently

You don’t need to install these locally — just open the chapter in the devcontainer and you’re ready to go.

## The Booking API

We start with **one RPC method**:

```proto
service BookingService {
  rpc CreateBooking(CreateBookingRequest) returns (CreateBookingResponse);
}

message CreateBookingRequest {
  string passenger = 1;
  string departure = 2;
  string destination = 3;
}

message CreateBookingResponse {
  string pnr = 1; // booking reference - passenger name record
  string status = 2;
}
```

**Message Flow**:

- client sends a booking request
- service responds with confirmation

## Build

From the chapter root:

1. Generate code:

```sh
make proto
```

2. Build the Booking service:

```sh
make services
```

3. Build the demo:

```sh
make demo
```

## Run

1. Start the **Booking gRPC Server**:

```sh
./build/services/booking
Starting gRPC listener on port :50051
Received booking request for John Doe (Vienna -> Berlin)
```

2. In a new terminal, start the client:

```sh
./build/demo/booking-grpc-client
Booking confirmed! PNR: PNR-12345, Status: CONFIRMED
```

## Next Steps

In the [next chapter](../02-connect-rpc) we will use **ConnectRPC** to bridge the gap between:

- **gRPC services**
- **HTTP/JSON clients** (cURL, browser, etc.)

...all using the **same protobuf service definition**.
