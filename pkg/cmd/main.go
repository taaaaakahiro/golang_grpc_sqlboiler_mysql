package main

import (
	"context"
	"golang-grpc-sqlboiler-mysql/pkg/server/models/command"
)

func main() {
	command.Run(context.Background())
}
