package grpcserver

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
	"log"
	"net"
	"nihago-user/database"
	"nihago-user/internal"
	"nihago-user/model"
	pb "nihago-user/pb/user"
)

type Server struct {
	db    *sql.DB
	cache *cache.Cache
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

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (s *Server) ConfigureCache() {
	s.cache = cache.New(cache.NoExpiration, cache.NoExpiration)
	companyList := internal.GetCompanies()
	for _, company := range companyList {
		s.cache.Set(string(company.Id), company.Name, cache.DefaultExpiration)
	}
	log.Println("Configuring company cache...")

}

func (s *Server) GetUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("Receive a GetUser request from client. User.Id: %d", in.Id)
	userModel := model.User{}
	UserModel := userModel.GetUserById(in.Id, s.db)
	User := s.convertUser(UserModel)
	return User, nil
}

func (s *Server) GetUserList(ctx context.Context, empty *pb.Empty) (*pb.UserList, error) {
	log.Printf("Receive a GetUsersList request from client")
	userModel := model.User{}
	users := userModel.GetUserList(s.db)
	var usersList []*pb.User

	for _, u := range users {
		usersList = append(usersList, s.convertUser(&u))
	}

	usersListPb := pb.UserList{UserList: usersList}
	return &usersListPb, nil
}

func (s *Server) convertUser(u *model.User) *pb.User {

	companyName, _ := s.cache.Get(string(u.CompanyId))
	if companyName == nil {
		companyName = "undefined"
	}

	return &pb.User{
		Id:          u.Id,
		Firstname:   u.Firstname,
		Lastname:    u.Lastname,
		Patronymic:  u.Patronymic,
		DateOfBirth: u.DateOfBirth,
		About:       u.About,
		Photo:       u.Photo,
		Company:     companyName.(string),
	}
}
