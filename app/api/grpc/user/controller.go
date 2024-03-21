package user

import (
	"context"
	"fmt"

	"user-service/app/helpers/log"
	"user-service/app/types"

	grpcApi "github.com/MicroservicesPractice/grpc-api/generated/user"
)

type UserGRPCServer struct {
	// type embedded to comply with Google lib
	grpcApi.UnimplementedUserServer
}

func (m *UserGRPCServer) CreateUser(ctx context.Context, request *grpcApi.CreateUserRequest) (*grpcApi.CreateUserResponse, error) {
	user := &types.User{
		Email:       request.Email,
		Password:    request.Password,
		Nickname:    request.Nickname,
		PhoneNumber: request.PhoneNumber,
	}

	if err := UserServiceInstance.CreateUser(user); err != nil {
		log.GrpcLog(log.Error, "user_service", fmt.Sprintf("can't create user: %v", err.Error()))
		return &grpcApi.CreateUserResponse{Message: "user was not created"}, nil
	}

	return &grpcApi.CreateUserResponse{Message: "user was created"}, nil
}

func (m *UserGRPCServer) GetUserPassword(ctx context.Context, request *grpcApi.GetUserPasswordRequest) (*grpcApi.GetUserPasswordResponse, error) {
	user, err := UserServiceInstance.GetUserPassword(request.Email)
	if err != nil {
		log.GrpcLog(log.Error, "user_service", fmt.Sprintf("can't get user password: %v", err.Error()))
		return &grpcApi.GetUserPasswordResponse{}, nil
	}

	return &grpcApi.GetUserPasswordResponse{Password: user.Password, Id: user.ID}, nil
}
