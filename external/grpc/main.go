package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
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

	fmt.Println("Please enter number(1:Unariy, 2:ClientStream, 3:BidirectionStream)")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	switch text {
	case "1":
		UnaryUser()
	case "2":
		ClientStreamUser()
	case "3":
		BidirectionStreamsUser()
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

func BidirectionStreamsUser() {
	stream, err := client.UserBidirectStream(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	sendNum := 4
	fmt.Printf("Please enter %d numbers.\n", sendNum)

	var sendEnd, recvEnd bool
	sendCount := 0
	for !(sendEnd && recvEnd) {
		// send
		if !sendEnd {
			scanner.Scan()
			inputNumber := scanner.Text()
			id, err := strconv.Atoi(inputNumber)
			if err != nil {
				log.Fatal(err)
			}

			sendCount++
			if err := stream.Send(&grpcpb.UserRequest{
				Id: int32(id),
			}); err != nil {
				fmt.Println(err)
				sendEnd = true
			}

			if sendCount == sendNum {
				sendEnd = true
				if err := stream.CloseSend(); err != nil {
					fmt.Println(err)
				}
			}
		}

		// receive
		if !recvEnd {
			if res, err := stream.Recv(); err != nil {
				if !errors.Is(err, io.EOF) {
					fmt.Println(err)
				}
				recvEnd = true
			} else {
				fmt.Printf("id:%d,name:%s,age:%d\n", res.GetId(), res.GetName(), res.GetAge())
			}
		}
	}
}
