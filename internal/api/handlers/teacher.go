package handlers

import (
	"context"

	
	"grpc_api/internal/models"
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

func (s *Server) GetTeachers(ctx context.Context, req *pb.GetTeachersRequest) (*pb.Teachers, error) {
	// err := req.Va()
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, err.Error())
	// }
	// Filtering, getting the filters from the request, another function
	filter, err := buildFilter(req.Teacher, &models.Teacher{})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// Sorting, getting the sort options from the request, another function
	sortOptions := buildSortOptions(req.GetSortBy())
	// Access the database to fetch data, another function

	teachers, err := mondodb.GetTeacherFromDB(ctx, sortOptions, filter)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Teachers{Teachers: teachers}, nil
}



