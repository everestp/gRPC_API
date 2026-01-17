package handlers

import (
	"context"
	"fmt"

	"grpc_api/internal/repositories/mondodb"
	pb "grpc_api/proto/gen"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddTeachers(ctx context.Context, req *pb.Teachers) (*pb.Teachers, error) {

	for _ , teacher := range req.GetTeachers(){
		if teacher.Id != ""{
			return nil, status.Error(codes.InvalidArgument, "Request is in  incorrect format -non-empty ID are not allowed")
		}
	}
 addedTeacher , err := mondodb.AddTeacherToDb(ctx, req.GetTeachers())
if err != nil {
	return nil, status.Error(codes.Internal, err.Error())
}

return &pb.Teachers{Teachers: addedTeacher}, nil


}

