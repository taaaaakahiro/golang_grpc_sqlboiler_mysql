run:
	go run ./pkg/server/main.go

sqlboiler:
	sqlboiler mysql -c sqlboiler.toml -o ./pkg/server/models --no-tests
	