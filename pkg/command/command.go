package command

import (
	"context"
	"database/sql"
	"fmt"
	"golang-grpc-sqlboiler-mysql/pkg/config"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"
)

func Run(ctx context.Context) {
	run(ctx)

}

func run(ctx context.Context) {
	cfg, err := config.LoadEnv(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// init mysql
	db, err := sql.Open("mysql", cfg.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		panic(err)
	}

	// 2. gRPCサーバーを作成
	s := grpc.NewServer()

	reflection.Register(s)
	// 3. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", cfg.Port)
		s.Serve(listener)
	}()

	// 4.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()

}
