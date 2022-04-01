package user

//
//import (
//	"context"
//	"log"
//	"nihago/service.user/internal"
//	"nihago/service.user/model"
//)
//
//type Server struct{}
//
//func (s *Server) GetUser(ctx context.Context, in *SingleUser) (*SingleUser, error) {
//	log.Printf("Receive a GetUser request from client. SingleUser.Id: %d", in.Id)
//	user := model.User{}
//	userModel := user.GetUserById(in.Id, s.db)
//	singleUser := convertUser(userModel)
//	return singleUser, nil
//}
//
//func (s *Server) GetUsers(ctx context.Context, empty *Empty) (*Users, error) {
//	panic("implement me")
//}
//
//func convertUser(u model.User) *SingleUser {
//	return &SingleUser{
//		Id:          u.Id,
//		Firstname:   u.Firstname,
//		Lastname:    u.Lastname,
//		Patronymic:  u.Patronymic,
//		DateOfBirth: u.DateOfBirth,
//		About:       u.About,
//		Photo:       u.Photo,
//		Company:     u.Company,
//	}
//}
