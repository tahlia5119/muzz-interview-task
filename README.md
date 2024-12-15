# Notes

Install protoc:
```bash
brew install protobuf
```

To generate golang grpc protobuf files:

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative muzz/explore-service.proto
```