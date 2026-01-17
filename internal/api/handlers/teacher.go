package handlers

import (
	"context"
	"fmt"

	"reflect"
	"strings"

	"grpc_api/internal/models"
	"grpc_api/internal/repositories/mondodb"
	pb "grpc_api/proto/gen"

	"go.mongodb.org/mongo-driver/bson"
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


func (s *Server) GetTeachers(ctx context.Context , req *pb.GetTeachersRequest)(*pb.Teachers , error){
	//Filtering getting the filters form the requeest another fucntion
    buildFilterForTeacher(req) 
	//Sorting gettting the sort options form the request another function
 buildSortOptions(req.GetSortBy())
	//Access the database to fetch data
	return nil, nil

}


func buildFilterForTeacher(req *pb.GetTeachersRequest){
	filter := bson.M{}

	var modelTeacher models.Teacher
modelVal := reflect.ValueOf(&modelTeacher).Elem()
modelType := modelVal.Type()

reqVal := reflect.ValueOf(req.Teacher).Elem()
	reqType := reqVal.Type()


	for i := 0 ; i < reqVal.NumField() ; i++{
		fieldVal := reqVal.Field(i)
		fieldName := reqType.Field(i).Name


		if fieldVal.IsValid() && !fieldVal.IsZero(){
			  modelFiled :=   modelVal.FieldByName(fieldName) 
			  if  modelFiled.IsValid() && modelFiled.CanSet(){
				modelFiled.Set(fieldVal)
			  }     
		}
                
	}

	//Now we iterate over the modelTeacher to build filter using bsom.M

	for i := 0 ; i < modelVal.NumField() ; i++{
		fieldVal := modelVal.Field(i)
		// fieldName := modelType.Field(i).Name
		if fieldVal.IsValid() && !fieldVal.IsZero(){
			bsonTag := modelType.Field(i).Tag.Get("bson")
			bsonTag = strings.TrimSuffix(bsonTag, ",omitempty")
			filter[bsonTag] = fieldVal.Interface().(string)
		}
	}
fmt.Println("Filter ==>",filter)
}


func buildSortOptions(sortFields []*pb.SortField) bson.D{
	var sortOptions bson.D
	for _ , sortField := range sortFields{
		order :=1
		if sortField.GetOrder() == pb.Order_DESC{
			order = -1
		}
		sortOptions = append(sortOptions, bson.E{Key: sortField.Field ,Value: order})
	}
	fmt.Println("Sort Options",sortOptions)

return sortOptions
}