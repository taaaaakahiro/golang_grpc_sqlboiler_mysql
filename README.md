# golang-grpc-sqlboiler-mysql

## Version
```
protoc: libprotoc 3.21.6
```

## Init Path
```
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

## Generate Protocol Buffer
```
$ cd ./pkg/proto
$ protoc --go_out=../grpc --go_opt=paths=source_relative \
	--go-grpc_out=../grpc --go-grpc_opt=paths=source_relative \
	main.proto
```

## brew install
 - grpcurl

## grpcurl ~ on runnning grpc server ~
```
$ grpcurl -plaintext localhost:8080 list // list service
$ grpcurl -plaintext localhost:8080 list grpcapp.UserService // list method
$ grpcurl -plaintext -d '{"id": "1"}' localhost:8080 grpcapp.UserService.User // request
```