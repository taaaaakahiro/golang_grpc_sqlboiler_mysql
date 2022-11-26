# golang-grpc-sqlboiler-mysql

## Procedure
1. Init Path
```
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

2. run server
```
$ make run
```

3. http request/grpcurl ~ on runnning grpc server ~
```
$ grpcurl -plaintext localhost:8080 list // list service
$ grpcurl -plaintext localhost:8080 list grpcapp.UserService // list method
```
```json
// request
$ grpcurl -plaintext -d '{"id": "1"}' localhost:8080 grpcapp.UserService.User

// response
{
  "id": 1,
  "name": "hoge",
  "age": 33
}
```

## Generate Protocol Buffer
```
$ cd ./pkg/proto
$ protoc --go_out=../grpc --go_opt=paths=source_relative \
	--go-grpc_out=../grpc --go-grpc_opt=paths=source_relative \
	user.proto
```

## Version
```
protoc: libprotoc 3.21.6
```

## brew install
 - grpcurl


