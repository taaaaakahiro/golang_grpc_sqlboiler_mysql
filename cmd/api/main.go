package main

import (
	"context"
	"golang-grpc-sqlboiler-mysql/pkg/command"
)

func main() {
	command.Run(context.Background())
}
