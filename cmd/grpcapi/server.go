package main

import (
	"fmt"
	"grpc_api/internal/api/handlers"
	"grpc_api/internal/repositories"
	pb "grpc_api/proto/gen"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file",err)
	}
	repositories.CreateMongoClient()
	s := grpc.NewServer()
  pb.RegisterExecsServiceServer(s, handlers.Server{})
   pb.RegisterStudentsServiceServer(s, handlers.Server{})
    pb.RegisterTeachersServiceServer(s, handlers.Server{})

	reflection.Register(s)
	port := os.Getenv("SERVER_PORT")
	fmt.Println("gRPC server is running at PORT ",port)
	list , err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error listing on  specified Port ",err)
	}

	err = s.Serve(list)
	if err != nil {
		log.Fatal("Failed to serve",err)
	}




}