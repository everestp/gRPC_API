package handlers

import (
	"context"
	"fmt"

	"grpc_api/internal/models"
	"grpc_api/internal/pkg/utils"
	"reflect"

	"grpc_api/internal/repositories/mondodb"
	pb "grpc_api/proto/gen"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) AddTeachers(ctx context.Context, req *pb.Teachers) (*pb.Teachers, error) {
client , err := 	mondodb.CreateMongoClient()
if err != nil {
	return nil, utils.ErrorHandler(err, "internal error")
}
defer client.Disconnect(ctx)
 newTeachers := make([]*models.Teacher , len(req.GetTeachers()))
for i, pbTeacher := range req.GetTeachers(){
	modelTeacher := models.Teacher{}
	pbVal := reflect.ValueOf(pbTeacher).Elem()
	modalVal := reflect.ValueOf(&modelTeacher).Elem()
	for i := 0; i < pbVal.NumField(); i++ {
		pbFiled :=pbVal.Field(i)
		fieldName := pbVal.Type().Field(i).Name
		modelField := modalVal.FieldByName(fieldName)
		
		 if modelField.IsValid() && modelField.CanSet(){
			modelField.Set(pbFiled)
		 } else{
			fmt.Printf("Field %s  is not valid or cannot be set\n",fieldName)
		 }
		
	}
   newTeachers[i]=&modelTeacher
}

 var addedTeacher []*pb.Teacher
for _ , teacher := range newTeachers{
	 result , err := client.Database("school").Collection("teachers").InsertOne(ctx,teacher)
	 if err != nil {
		return nil, utils.ErrorHandler(err, "Error  adding the  value to database")
	 }

objectId  , ok  := 	  result.InsertedID.(primitive.ObjectID)
if ok {
	teacher.Id = objectId.Hex()
}
pbTeacher := &pb.Teacher{}
modalVal := reflect.ValueOf(*teacher)
pbVal := reflect.ValueOf(pbTeacher).Elem()
 for i := 0; i < modalVal.NumField(); i++{
	modelField:= modalVal.Field(i)
	modelFiledType := modalVal.Type().Field(i)
	// pbFieldType := pbVal.Type().Field(i)
	pbField := pbVal.FieldByName(modelFiledType.Name)
	if pbField.IsValid() && pbField.CanSet(){
		pbField.Set(modelField)
	}
 }
 addedTeacher = append(addedTeacher, pbTeacher)
}


fmt.Println("NEw Teacher",newTeachers)
return &pb.Teachers{Teachers: addedTeacher}, nil


}