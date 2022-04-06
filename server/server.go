package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Sakamoto0525/gRPC-Tutorial/gen/api"
	"github.com/Sakamoto0525/gRPC-Tutorial/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 8000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	api.RegisterPancakeBakerServiceServer(
		server,                    // gRPCサーバー
		handler.NewBakerHandler(), // リクエストを処理するハンドラ
	)
	reflection.Register(server)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}
