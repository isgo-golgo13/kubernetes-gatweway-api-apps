## Go gRPC App (for Gateway API using GRPCRoute CR)

This is a Go (1.23) gRPC client and server application as a source for downstream use with the Kubernetes Gateway API and the specific use of the Gateway API CR `GRPCRoute`.


### The gRPC Proto Source

The following .proto  source file in directory `proto/send.proto` provides the following API methods

- send w/ byte.Buffer
- send w/ timeout and byte.Buffer
- send all w/ streaming byte.Buffer

```proto
syntax = "proto3";

package send;

option go_package = "./proto;proto";

// The send service definition.
service SendService {
  // Sends data with ID and bytes.
  rpc Send (SendRequest) returns (SendResponse) {}

  // Sends data with ID, bytes and a timeout.
  rpc SendWithTimeout (SendWithTimeoutRequest) returns (SendResponse) {}

  // Streaming version of Send.
  rpc SendAll (stream SendRequest) returns (stream SendResponse) {}
}

// The request message containing the ID and bytes.
message SendRequest {
  int32 id = 1;
  bytes data = 2;
}

// The request message containing the ID, bytes and timeout.
message SendWithTimeoutRequest {
  int32 id = 1;
  bytes data = 2;
  int32 timeout = 3;
}

// The response message containing the number of bytes sent.
message SendResponse {
  int32 bytes_sent = 1;
}
```


### Generating the gRPC Proto Sources to Go

```shell
protoc --go_out=proto --go-grpc_out=proto proto/send.proto
```

- The **--go_out=proto**: This flag tells protoc to generate the Go files in the proto directory.
- The **--go-grpc_out=proto**: This flag tells protoc to generate the gRPC-specific Go files in the proto directory.
- The **proto/send.proto**: The path to the .proto file

### Client and Server Configuration

The client.go and server.go sources reference external configuration of ports and address using the provided `.env` configuration file and the .

