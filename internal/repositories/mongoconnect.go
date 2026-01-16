package repositories

import (
	"context"
	
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
		log.Println("Error connecting to Mongodb",err)

	}
	err = client.Ping(ctx, nil)
	if err != nil{
		log.Println("Error connecting to Mongodb",err)
		return nil, err

	}
   log.Println("Connected to mongodb")
   return  client , nil
}