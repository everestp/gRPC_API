package handlers

import pb "grpc_api/proto/gen"


type Server struct{
	pb.UnimplementedExecsServiceServer
	pb.UnimplementedStudentsServiceServer
	pb.UnimplementedTeachersServiceServer
}




