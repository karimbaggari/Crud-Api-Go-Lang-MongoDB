package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"scriptology/app/controllers/User/controller"
	"scriptology/app/controllers/User/service" 
)

var (
	server        *gin.Engine
	userService   service.UserService
	userController controller.UserController
	ctx           context.Context
	userCollection *mongo.Collection
	mongoClient   *mongo.Client
	err           error
)

func init() {
	fmt.Println("server started")
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		fmt.Println("error cut")
		log.Fatal(err)
	}
	fmt.Println("connection with db started")

	err = mongoClient.Ping(ctx, readpref.Primary())
	fmt.Println("error may be cut")
	if err != nil {
		fmt.Println("error after Ping:", err) // Add this line to log the error
		log.Fatal(err)
	}
	fmt.Println("mongo connection established")
	userCollection = mongoClient.Database("userdb").Collection("users")
	userService = service.NewUserService(userCollection, ctx)
	userController = controller.NewUserController(userService)
	server = gin.Default()
	fmt.Println("connection fully passed")
}

func main() {
	fmt.Println("server started")
	defer mongoClient.Disconnect(ctx)

	basepath := server.Group("/v1")
	userController.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":9090"))
	fmt.Println("server started")
}
