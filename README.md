Tinkoff invest API gRPC [go](https://go.dev) example.

## Generate gRPC client
1. Update [./protos/](./protos/) from https://github.com/Tinkoff/invest-python/tree/main/protos/ (if required)
1. Install `protoc` version 3 https://grpc.io/docs/protoc-installation/
1. Install plugins
    ```sh
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```
1. Generate code
    ```sh
    mkdir -p investapi && \
    protoc --go-grpc_out=investapi/ --go_out=investapi/ --proto_path=protos/ protos/tinkoff/invest/grpc/*.proto
    ```

## Build
```sh
go build .
```

## Usage
```sh
./tinkoff-invest-example -t your-api-token
```

## Documentation
https://tinkoff.github.io/investAPI/
