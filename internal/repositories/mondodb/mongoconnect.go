package mondodb

import (
	"context"
	"grpc_api/internal/pkg/utils"

	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func CreateMongoClient()(*mongo.Client , error){
	ctx := context.Background()
  url := os.Getenv("MONGODB_URI")
	 client , err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil{
	 return nil, utils.ErrorHandler(err, "Unable to connect to MongoDB")

	}
	err = client.Ping(ctx, nil)
	if err != nil{
		
		return nil,  utils.ErrorHandler(err, "Unable to ping database")

	}
   log.Println("Connected to mongodb")
   return  client , nil
}