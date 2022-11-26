# golang-grpc-sqlboiler-mysql

## Version
```
protoc: libprotoc 3.21.6
```

## Generate Protocol Buffer
```
$ cd ./pkg/proto
$ protoc --go_out=../grpc --go_opt=paths=source_relative \
	--go-grpc_out=../grpc --go-grpc_opt=paths=source_relative \
	main.proto
```