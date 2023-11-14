package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "lecture_18/proto/gw"
	"log"
	"net"
)

func main() {
	fmt.Println("server")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	followServer := NewServer()

	pb.RegisterUserServiceServer(grpcServer, followServer)

	log.Fatalln(grpcServer.Serve(listener))
}
