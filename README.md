Tinkoff invest API gRPC example

## Generate gRPC client
1. Update [./protos/](./protos/) from https://github.com/Tinkoff/invest-python/tree/main/protos/ if required
1. Install `protoc` version 3 (https://developers.google.com/protocol-buffers)
1. Install `go` plugin for `protoc`
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```
1. Install `go-grpc` plugin for `protoc`
```sh
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
1. Generate
```sh
protoc protos/tinkoff/invest/grpc/* --go-grpc_out=investapi/ --go_out investapi/ -I protos/
```

## Build
```sh
go build .
```

## Usefull links
- https://grpc.io/docs/languages/go/quickstart/
- https://github.com/Tinkoff/invest-python
- https://tinkoff.github.io/investAPI/