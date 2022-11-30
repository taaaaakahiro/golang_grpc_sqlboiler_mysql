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
// 1. Unary
$ grpcurl -plaintext -d '{"id": 1}' localhost:8080 grpcapp.UserService.User
// res
{
  "id": 1,
  "name": "user1",
  "age": 11
}
// 2. Server Streaming
$ grpcurl -plaintext -d '{"id": 1}' localhost:8080 grpcapp.UserService.UserServerStream
// res
{
  "id": 1,
  "name": "user1",
  "age": 11
}
{
  "id": 2,
  "name": "user2",
  "age": 22
}
{
  "id": 3,
  "name": "user3",
  "age": 33
}
{
  "id": 4,
  "name": "user4",
  "age": 44
}
// 3. Client Streaming
$ grpcurl -plaintext -d '{"id": 1}{"id": 2}' localhost:8080 grpcapp.UserService.UserClientStream
// res
{
  "user": [
    {
      "id": 1,
      "name": "user1",
      "age": 11
    },
    {
      "id": 2,
      "name": "user2",
      "age": 22
    }
  ]
}
// 4. Bidirectional Streaming
$ grpcurl -plaintext -d '{"id": 1}{"id": 2}' localhost:8080 grpcapp.UserService.UserBidirectStream
// res
{
  "id": 1,
  "name": "user1",
  "age": 11
}
{
  "id": 2,
  "name": "user2",
  "age": 22
}
```

4. http request from grpc client
	1. run
		```
		$ cd .
		$ make client
		```
	2. input number(1:Unary, 2:ClientStream,3:BidirecrionStream)
	3. input four numbers(id) & enter
	 - #2 ClientStream
		```
		// example input #1~4
		[
			id:1 name:"user1" age:11
			id:2 name:"user2" age:22
			id:3 name:"user3" age:33
			id:4 name:"user4" age:44
		]
		```
	 - #3 BidirecrionStream
		```
		1
		id:1,name:user1,age:11
		2
		id:2,name:user2,age:22
		3
		id:3,name:user3,age:33
		4
		id:4,name:user4,age:44
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


