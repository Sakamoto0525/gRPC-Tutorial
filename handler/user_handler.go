package handler

import (
	"context"
	"time"

	"github.com/Sakamoto0525/gRPC-Tutorial/gen/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (uh *UserHandler) Index(
	ctx context.Context,
	req *api.UserRequest,
) (*api.UserResponse, error) {
	now := time.Now()

	return &api.UserResponse{
		User: &api.User{
			Id:      1,
			Name:    "田中太郎",
			Age:     21,
			Tel:     123456789,
			Address: "example@test.com",
			CreateTime: &timestamppb.Timestamp{
				Seconds: now.Unix(),
				Nanos:   int32(now.Nanosecond()),
			},
		},
	}, nil
}

func (uh *UserHandler) Show(
	ctx context.Context,
	req *api.UserRequest,
) (*api.UserResponse, error) {
	now := time.Now()

	return &api.UserResponse{
		User: &api.User{
			Id:      1,
			Name:    "田中太郎",
			Age:     21,
			Tel:     123456789,
			Address: "example@test.com",
			CreateTime: &timestamppb.Timestamp{
				Seconds: now.Unix(),
				Nanos:   int32(now.Nanosecond()),
			},
		},
	}, nil
}
