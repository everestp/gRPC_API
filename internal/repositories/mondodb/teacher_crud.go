package mondodb

import (
	"context"
	"fmt"

	"grpc_api/internal/models"
	"grpc_api/internal/pkg/utils"
	"reflect"


	pb "grpc_api/proto/gen"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func AddTeacherToDb(ctx context.Context, teachersFormReq []*pb.Teacher) ( []*pb.Teacher,  error) {
	client, err := CreateMongoClient()
	if err != nil {
		return nil,utils.ErrorHandler(err, "internal error")
	}
	defer client.Disconnect(ctx)
	newTeachers := make([]*models.Teacher, len(teachersFormReq)
	for i, pbTeacher := range teachersFormReq {
		newTeachers[i] = mapPbTeacherToModelTeacher(pbTeacher)

	}

	var addedTeacher []*pb.Teacher
	for _, teacher := range newTeachers {
		result, err := client.Database("school").Collection("teachers").InsertOne(ctx, teacher)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error  adding the  value to database")
		}

		ObjectID, ok := result.InsertedID.(primitive.ObjectID)
		if ok {
			teacher.Id = ObjectID.Hex()
		}
		pbTeacher := mapModelTeacherToPb(teacher)
		addedTeacher = append(addedTeacher, pbTeacher)
	}
	return  addedTeacher, nil
}

func mapModelTeacherToPb(teacher *models.Teacher) *pb.Teacher {
	pbTeacher := &pb.Teacher{}
	modalVal := reflect.ValueOf(*teacher)
	pbVal := reflect.ValueOf(pbTeacher).Elem()
	for i := 0; i < modalVal.NumField(); i++ {
		modelField := modalVal.Field(i)
		modelFiledType := modalVal.Type().Field(i)
		// pbFieldType := pbVal.Type().Field(i)
		pbField := pbVal.FieldByName(modelFiledType.Name)
		if pbField.IsValid() && pbField.CanSet() {
			pbField.Set(modelField)
		}
	}
	return pbTeacher
}

func mapPbTeacherToModelTeacher(pbTeacher *pb.Teacher) *models.Teacher{
	modelTeacher := models.Teacher{}
	pbVal := reflect.ValueOf(pbTeacher).Elem()
	modalVal := reflect.ValueOf(&modelTeacher).Elem()
	for i := 0; i < pbVal.NumField(); i++ {
		pbFiled := pbVal.Field(i)
		fieldName := pbVal.Type().Field(i).Name
		modelField := modalVal.FieldByName(fieldName)

		if modelField.IsValid() && modelField.CanSet() {
			modelField.Set(pbFiled)
		} else {
			fmt.Printf("Field %s  is not valid or cannot be set\n", fieldName)
		}

	}
	 return &modelTeacher
}