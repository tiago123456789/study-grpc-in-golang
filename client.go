package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/tiago123456789/study-grpc/pb"
	"google.golang.org/grpc"
)

func main() {

	connect, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Could not connect to gRPC server: %v", err)
	}

	defer connect.Close()

	client := pb.NewUserServiceClient(connect)
	AddUserStreamBoth(client)
	// AddUsers(client)
	// AddUser(client)
	// AddUserStream(client)
}

func AddUser(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "teste",
		Email: "teste@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatal("Could not make to gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserStream(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "teste",
		Email: "teste@gmail.com",
	}

	res, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatal("Could not make to gRPC request: %v", err)
	}

	for {
		stream, err := res.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Could not receive the message: %v", err)
		}

		fmt.Println("Status:", stream.GetStatus())

	}

}

func AddUsers(client pb.UserServiceClient) {
	users := []*pb.User{
		&pb.User{
			Id:    "t1",
			Name:  "t 1",
			Email: "t1@gmail.com",
		},
		&pb.User{
			Id:    "t2",
			Name:  "t 2",
			Email: "t2@gmail.com",
		},
		&pb.User{
			Id:    "t3",
			Name:  "t 3",
			Email: "t3@gmail.com",
		},
		&pb.User{
			Id:    "t4",
			Name:  "t 4",
			Email: "t4@gmail.com",
		},
		&pb.User{
			Id:    "t5",
			Name:  "t 5",
			Email: "t5@gmail.com",
		},
		&pb.User{
			Id:    "t6",
			Name:  "t 6",
			Email: "t6@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatal("Error creating request: %v", err)
	}

	for _, req := range users {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Error creating response: %v", err)
	}

	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	users := []*pb.User{
		&pb.User{
			Id:    "t1",
			Name:  "t 1",
			Email: "t1@gmail.com",
		},
		&pb.User{
			Id:    "t2",
			Name:  "t 2",
			Email: "t2@gmail.com",
		},
		&pb.User{
			Id:    "t3",
			Name:  "t 3",
			Email: "t3@gmail.com",
		},
		&pb.User{
			Id:    "t4",
			Name:  "t 4",
			Email: "t4@gmail.com",
		},
		&pb.User{
			Id:    "t5",
			Name:  "t 5",
			Email: "t5@gmail.com",
		},
		&pb.User{
			Id:    "t6",
			Name:  "t 6",
			Email: "t6@gmail.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range users {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
			}

			fmt.Printf("Receiving user %v with status %v", res, res.GetStatus())

		}
		close(wait)
	}()
	<-wait
}
