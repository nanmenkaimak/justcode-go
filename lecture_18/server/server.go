package main

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "lecture_18/proto/gw"
	"time"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	followers map[string][]User
}

func NewServer() Server {
	return Server{
		followers: users,
	}
}

func (s Server) GetAllFollowers(request *pb.GetAllFollowersRequest, stream pb.UserService_GetAllFollowersServer) error {
	follow := s.followers[request.GetUserId()]

	for _, user := range follow {
		if err := stream.Send(&pb.GetAllFollowersResponse{
			User: &pb.User{
				Id:          user.ID,
				FirstName:   user.FirstName,
				LastName:    user.LastName,
				Username:    user.Username,
				Email:       user.Email,
				Password:    user.Password,
				IsConfirmed: user.IsConfirmed,
			},
		}); err != nil {
			return status.Errorf(codes.Internal, "fetch: unexpected stream: %v", err)
		}
		time.Sleep(time.Duration(5) * time.Second)
	}
	return nil
}

type User struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsConfirmed bool   `json:"is_confirmed"`
}

var users = map[string][]User{
	"1": {
		{
			ID:          "1",
			FirstName:   "a",
			LastName:    "a",
			Username:    "a",
			Email:       "a",
			Password:    "a",
			IsConfirmed: true,
		},
		{
			ID:          "2",
			FirstName:   "b",
			LastName:    "b",
			Username:    "b",
			Email:       "b",
			Password:    "b",
			IsConfirmed: true,
		},
	},
}
