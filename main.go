package main

import (
	"context"
	"fmt"
	"go-crud/controllers"
	"go-crud/services"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroller controllers.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()

	mongoConnection := options.Client().ApplyURI("mongodb+srv://sagardolia:JfYTkgsJ52e9ArIe@dashboard.a7h1nux.mongodb.net/?retryWrites=true&w=majority")

	mongoclient, err = mongo.Connect(ctx, mongoConnection)
	if err != nil {
		log.Fatal(err)
	}

	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo Connection is established")

	usercollection = mongoclient.Database("userDb").Collection("users")

	userservice = services.NewUserService(usercollection, ctx)

	usercontroller = controllers.New(userservice)

	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("api/v1")
	usercontroller.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))
}
