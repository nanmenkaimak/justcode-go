package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, _ := grpc.Dial(":50051", opts...)

	client := NewClient(conn)
	if err := client.GetAllFollowers(context.Background(), "1"); err != nil {
		log.Fatal(err)
	}
}
