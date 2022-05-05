Tinkoff invest API gRPC [go](https://go.dev) example.

## Usage
```sh
./tinkoff-invest-example -t your-api-token
```

## Build
```sh
go build .
```

## Generate gRPC client
1. Install `protoc` version 3 https://grpc.io/docs/protoc-installation/
1. Install plugins
    ```sh
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```
1. Generate code
    ```sh
    mkdir -p tinkoff/investapi/ && \
    protoc --go-grpc_out=tinkoff/investapi/ --go_out=tinkoff/investapi/ --proto_path=protos/ protos/*.proto
    ```

## Update protocol buffer files
https://github.com/Tinkoff/investAPI/tree/main/src/docs/contracts

## Documentation
https://tinkoff.github.io/investAPI/
