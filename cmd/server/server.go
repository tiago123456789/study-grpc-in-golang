package main

import (
	"fmt"
	"log"
	"net"

	"github.com/tiago123456789/study-grpc/pb"
	"github.com/tiago123456789/study-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("Could not connect in server in address: localhost:50051", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)
	fmt.Println("Server is running in address: localhost:50051")

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatal("Could not connect in server in address: localhost:50051", err)
	}

}
