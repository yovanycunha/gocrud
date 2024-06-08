package main

import (
	"context"
	"fmt"
	"go-crud/initializers"
	"go-crud/pkg/controllers"
	"go-crud/pkg/services"
	"log"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server *gin.Engine
	mongoclient *mongo.Client
	ctx context.Context
	userCollection *mongo.Collection
	userService services.UserService
	userController controllers.UserController
)

func init(){
	initializers.LoadEnv()

	ctx := context.TODO()

	mongoConn := options.Client().ApplyURI("mongodb://localhost:27017/")
	mongoclient, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal(err)
	}

	err = mongoclient.Ping(ctx,readpref.Primary())
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection successfully established!")

	userCollection = mongoclient.Database("gocruddb").Collection("users")
	userService = services.NewUserService(userCollection,ctx)
	userController = controllers.New(userService)

	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)
	
	basepath := server.Group("/api")
	userController.RegisterUserRoutes(basepath)

	log.Fatal(server.Run())

}

// func homeRoute(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "Chama!",
// 	})
// }