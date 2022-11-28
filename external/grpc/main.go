package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	grpcpb "golang-grpc-sqlboiler-mysql/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	scanner *bufio.Scanner
	client  grpcpb.UserServiceClient
)

func main() {
	fmt.Println("start gRPC Client.")

	// 1. 標準入力から文字列を受け取るスキャナを用意
	scanner = bufio.NewScanner(os.Stdin)

	// 2. gRPCサーバーとのコネクションを確立
	p := os.Getenv("PORT")
	host := fmt.Sprintf("localhost:%s", p)
	conn, err := grpc.Dial(
		host,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// 3. gRPCクライアントを生成
	client = grpcpb.NewUserServiceClient(conn)
	User()
}

func User() {
	fmt.Println("Please enter userID(#1~4)")
	scanner.Scan()
	text := scanner.Text()
	id, _ := strconv.Atoi(text)

	req := &grpcpb.UserRequest{

		Id: int32(id),
	}
	res, err := client.User(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("ID:%d, Name:%s, Age:%d", res.Id, res.Name, res.Age)
	}
}
