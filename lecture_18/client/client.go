package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	pb "lecture_18/proto/gw"
	"log"
)

type Client struct {
	client pb.UserServiceClient
}

func NewClient(conn grpc.ClientConnInterface) Client {
	return Client{
		client: pb.NewUserServiceClient(conn),
	}
}

func (c Client) GetAllFollowers(ctx context.Context, userID string) error {
	stream, err := c.client.GetAllFollowers(ctx, &pb.GetAllFollowersRequest{
		UserId: userID,
	})
	if err != nil {
		return err
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("Finished")
				return nil
			}
			return err
		}
		followers := res.GetUser()
		fmt.Println(followers)
	}
}
