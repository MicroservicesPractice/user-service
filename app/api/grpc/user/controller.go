package user

import (
	"context"
	"log"
)

type UserGRPCServer struct {
	// type embedded to comply with Google lib
	UnimplementedUserServer
}

func (m *UserGRPCServer) CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {
	log.Println("CreateUser called")
	return &CreateUserResponse{Message: "dsfsdf"}, nil
}

func (m *UserGRPCServer) GetUserPassword(ctx context.Context, request *GetUserPasswordRequest) (*GetUserPasswordResponse, error) {
	log.Println("GetUserPassword called")
	return &GetUserPasswordResponse{Password: "dsfsdf"}, nil
}
