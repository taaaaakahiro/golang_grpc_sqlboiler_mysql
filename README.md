# golang-grpc-sqlboiler-mysql

## Procedure
- 1 → 2 → 3 or 4
1. init path $ set environment
```
$ export PATH="$PATH:$(go env GOPATH)/bin"
$ export PORT=< grpc server port >
$ export MYSQL_DSN=< mysql dsn >

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
$ grpcurl -plaintext -d '{"id": 1}' localhost:8080 grpcapp.UserService.User
// response
{
  "id": 1,
  "name": "user1",
  "age": 11
}
```

4. http request from grpc client
```
$ cd .
$ make client
```

## regenerate by Protocol Buffer
```
$ cd ./pkg/proto
$ protoc --go_out=../grpc --go_opt=paths=source_relative \
	--go-grpc_out=../grpc --go-grpc_opt=paths=source_relative \
	user.proto
```

## regenerate by sqlboiler
```
$ cd .
$ make sqlboiler
```

## Version
```
protoc: libprotoc 3.21.6
```

## brew install
 - grpcurl


