package grpcserver

import (
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"nihago-users/database"
	"nihago-users/model"
	"nihago-users/pb/user"

	_ "github.com/lib/pq"
)

type Server struct {
	db *sql.DB
}

func (s *Server) Run() {
	fmt.Printf("Starting a GRPC server at :8090")

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initialize User database

	pgsql := database.PostgreSQL{}
	s.db = pgsql.Init("postgres", "1q2w3e4r", "postgres", "15432")
	defer pgsql.Instance.Close()

	grpcServer := grpc.NewServer()

	user.RegisterUserServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (s *Server) GetUser(ctx context.Context, in *user.SingleUser) (*user.SingleUser, error) {
	log.Printf("Receive a GetUser request from client. SingleUser.Id: %d", in.Id)
	userModel := model.User{}
	singleUserModel := userModel.GetUserById(in.Id, s.db)
	singleUser := convertUser(singleUserModel)
	return singleUser, nil
}

func (s *Server) GetUsers(ctx context.Context, empty *user.Empty) (*user.Users, error) {
	panic("implement me")
}

func convertUser(u *model.User) *user.SingleUser {
	return &user.SingleUser{
		Id:          u.Id,
		Firstname:   u.Firstname,
		Lastname:    u.Lastname,
		Patronymic:  u.Patronymic,
		DateOfBirth: u.DateOfBirth,
		About:       u.About,
		Photo:       u.Photo,
		Company:     u.Company,
	}
}
