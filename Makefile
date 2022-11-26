run:
	go run ./pkg/cmd/main.go

sqlboiler:
	sqlboiler mysql -c sqlboiler.toml -o ./pkg/server/models --no-tests
	