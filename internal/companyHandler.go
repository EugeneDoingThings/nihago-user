package internal

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "nihago-user/pb/company"
)

//type Company struct {
//	id   int32
//	name string
//}

func GetCompanies() []*pb.Company {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8100", grpc.WithInsecure())
	if err != nil {
		err = fmt.Errorf("did not connect: %s", err)
		panic(err)
	}

	defer conn.Close()

	u := pb.NewCompanyServiceClient(conn)

	//TODO зарефакторить
	empty := &pb.Empty{}

	companyList, err := u.GetCompanyList(context.Background(), empty)
	if err != nil {
		err = fmt.Errorf("error when calling GetCompany: %s", err)
		panic(err)
	}

	//list := make(map[int32]string)
	//
	//for _, company := range companyList.CompanyList {
	//	list[company.Id] = company.Name
	//}

	return companyList.CompanyList
}
