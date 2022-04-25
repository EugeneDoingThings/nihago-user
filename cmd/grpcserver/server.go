package grpcserver

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	"nihago-user/database"
	"nihago-user/internal"
	"nihago-user/model"
	pb "nihago-user/pb/user"
)

type Server struct {
	db *sql.DB
}

func (s *Server) Run() {
	fmt.Printf("Starting a GRPC server at :8090 \n")

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initialize User database

	pgsql := database.PostgreSQL{}
	s.db = pgsql.Init("postgres", "1q2w3e4r", "postgres", "15432")
	defer pgsql.Instance.Close()

	//TODO Восстановление после паники
	//grpcServer := grpc.NewServer(
	//	grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
	//		grpc_recovery.StreamServerInterceptor(),
	//	)),
	//	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
	//		grpc_recovery.UnaryServerInterceptor(),
	//	)),
	//)
	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (s *Server) GetUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("Receive a GetUser request from client. User.Id: %d", in.Id)
	userModel := model.User{}
	UserModel := userModel.GetUserById(in.Id, s.db)
	User := convertUser(UserModel)
	return User, nil
}

func (s *Server) GetUserList(ctx context.Context, empty *pb.Empty) (*pb.UserList, error) {
	log.Printf("Receive a GetUsersList request from client")
	userModel := model.User{}
	users := userModel.GetUserList(s.db)
	var usersList []*pb.User

	for _, u := range users {
		usersList = append(usersList, convertUser(&u))
	}

	companyList := internal.GetCompanies()
	log.Println(companyList)

	usersListPb := pb.UserList{UserList: usersList}
	return &usersListPb, nil
}

func convertUser(u *model.User) *pb.User {
	return &pb.User{
		Id:          u.Id,
		Firstname:   u.Firstname,
		Lastname:    u.Lastname,
		Patronymic:  u.Patronymic,
		DateOfBirth: u.DateOfBirth,
		About:       u.About,
		Photo:       u.Photo,
		CompanyId:   u.CompanyId,
	}
}
