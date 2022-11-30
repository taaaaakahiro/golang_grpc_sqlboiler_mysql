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

	// 1. connect to gRPC server
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

	// 2. generate gRPC client
	client = grpcpb.NewUserServiceClient(conn)

	fmt.Println("Please enter number(1:Unariy, 2:ClientStream)")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	switch text {
	case "1":
		UnaryUser()
	case "2":
		ClientStreamUser()

	}

}

func UnaryUser() {
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

func ClientStreamUser() {
	stream, err := client.UserClientStream(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	sendCount := 4
	fmt.Printf("Please input %d numbers(1~4).\n", sendCount)
	for i := 0; i < sendCount; i++ {
		scanner.Scan()
		inputNumber := scanner.Text()
		id, err := strconv.Atoi(inputNumber)
		if err != nil {
			fmt.Println("please input number.")
			continue
		}

		if err := stream.Send(&grpcpb.UserRequest{
			Id: int32(id),
		}); err != nil {
			fmt.Println(err)
			return
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetUser())
	}
}
