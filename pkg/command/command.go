package command

import (
	"context"
	"database/sql"
	"fmt"
	"golang-grpc-sqlboiler-mysql/pkg/config"
	grpcpb "golang-grpc-sqlboiler-mysql/pkg/grpc"
	"golang-grpc-sqlboiler-mysql/pkg/infrastracture/persistence"
	"golang-grpc-sqlboiler-mysql/pkg/server"

	"log"
	"net"
	"os"
	"os/signal"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"
)

const (
	exitOK  = 0
	exitErr = 1
)

func Run(ctx context.Context) {
	os.Exit(run(ctx))
}

func run(ctx context.Context) int {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup logger: %s\n", err)
		return exitErr
	}
	defer logger.Sync()

	cfg, err := config.LoadEnv(ctx)
	if err != nil {
		logger.Error("failed to load env", zap.Error(err))
		return exitErr
	}
	// init mysql
	db, err := sql.Open("mysql", cfg.Dsn)
	if err != nil {
		logger.Error("failed to open mysql", zap.Error(err))
		return exitErr
	}
	err = db.PingContext(ctx)
	if err != nil {
		logger.Error("failed to ping db", zap.Error(err))
		return exitErr
	}

	repositories, err := persistence.NewRepositories(ctx, db)
	if err != nil {
		logger.Error("failed to init repository", zap.Error(err))
		return exitErr
	}

	// init grpc
	s := grpc.NewServer()

	// init server
	servers, err := server.NewServer(repositories)
	if err != nil {
		logger.Error("failed to init server", zap.Error(err))
		return exitErr
	}

	// register gRPC method
	grpcpb.RegisterGreetingServiceServer(s, servers)

	reflection.Register(s)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		logger.Error("failed to listen", zap.Error(err))
		return exitErr
	}

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

	return exitOK
}
