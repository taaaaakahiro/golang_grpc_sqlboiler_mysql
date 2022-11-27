run:
	go run ./cmd/api/main.go

sqlboiler:
	sqlboiler mysql -c sqlboiler.toml -o ./pkg/models --no-tests
	
client:
	go run ./external/grpc/main.go