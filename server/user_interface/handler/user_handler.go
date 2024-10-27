package handler

import (
	pb "com.graffity/mission-sample/pkg/grpc"
	"com.graffity/mission-sample/server/applicationservice/usecase"
	"context"
	"log"
)

type UserHandler struct {
	pb.UnimplementedUsersServiceServer
	usecase *usecase.UserUsecase
}

func NewUserHandler(usecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (u *UserHandler) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// service経由でrepository行ってuserとtoken取得
	user, results, err := u.usecase.Save(ctx)
	if err != nil {
		return nil, err
	}

	pbUser := pb.User{
		Id: uint64(user.ID),
	}

	log.Println("clear missions", results)
	return &pb.CreateUserResponse{
		User: &pbUser,
	}, nil
}
