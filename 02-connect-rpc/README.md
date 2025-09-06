# 02-connect-rpc

In this chapter, we extend the minimal gRPC client/server from Chapter 1 by integrating [ConnectRPC](https://connectrpc.com/).

ConnectRPC allows us to expose the same protobuf-defined `BookingService` over **gRPC**, **gRPC-Web**, and **HTTP/JSON** simultaneously.

**Why this matters**:

In real-world systems, **gRPC endpoints are usually not exposed directly to customer-facing clients** (like browsers, mobile apps, or third-party developers). Instead, customer-facing clients typically interact over **HTTP/JSON** via an API gateway, while **internal services** communicate with each other over gRPC.

This pattern lets us:

- Keep efficient, strongly-typed **gRPC** between microservices.
- Offer developer-friendly, widely-compatible **HTTP/JSON APIs** to external clients.

With ConnectRPC, we get both worlds **from the same protobuf contract**.

## Tools

Same as Chapter 1, plus:

- **protoc-gen-connect-go** – generates ConnectRPC service stubs

(Already pre-installed in the devcontainer for this chapter.)

Update `buf.gen.yaml` to now build with `protoc-gen-connect-go` instead of `protoc-gen-go-grpc`:

```yaml
plugins:
  - local: protoc-gen-connect-go
    out: gen/go
    opt: paths=source_relative 
```

## The Booking API

No changes from Chapter 1 — we still have one RPC:

```proto
service BookingService {
  rpc CreateBooking(CreateBookingRequest) returns (CreateBookingResponse);
}
```

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

1. Start the **Booking service** with Connect RPC:

```sh
./build/services/booking
```

2. In a new terminal, start the **grpc client**:

```sh
./build/demo/booking-grpc-client
Booking confirmed! PNR: PNR-12345, Status: CONFIRMED
```

### Test via plain HTTP/JSON

Send a booking request directly with cURL:

```sh
./build/demo/booking-http-client | jq
{
  "pnr": "PNR-12345",
  "status": "CONFIRMED"
}
```
