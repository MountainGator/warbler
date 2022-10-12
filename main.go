package main

import (
	"context"
	"log"

	"github.com/MountainGator/warbler/controllers"
	"github.com/MountainGator/warbler/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	us          services.UserService
	ws          services.WarbleService
	uc          controllers.UserController
	user_coll   *mongo.Collection
	warble_coll *mongo.Collection
	client      *mongo.Client
	err         error
	key         []byte
	store       *sessions.CookieStore
)

func init() {
	// if err = godotenv.Load(); err != nil {
	// 	log.Println("No .env file found")
	// }

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("error connecting to mongo", err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal("error pinging mongo", err)
	}

	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)

	store.Options.HttpOnly = false
	store.Options.Secure = false

	user_coll = client.Database("playlist_db").Collection("users")
	warble_coll = client.Database("playlist_db").Collection("warbles")
	us = services.NewUserService(user_coll, store, context.TODO())
	ws = services.NewWarbleService(warble_coll, context.TODO())
	uc = controllers.NewUserController(us, ws, store)
}

func main() {
	r := gin.Default()
	user_router := r.Group("/user", uc.Auth)

	r.POST("/login", uc.UserLogin)
	r.POST("/new-user", uc.CreateUser)
	r.GET("/all-warbles", uc.FindAll)

	user_router.GET("/logout", uc.Logout)
	user_router.GET("/user-id", uc.GetUserDetails)
	user_router.PATCH("/update", uc.UpdateUser)
	user_router.DELETE("/delete", uc.DeleteUser)

	user_router.POST("/new-warble", uc.CreateWarble)

	user_router.PATCH("/edit-warble", uc.EditWarble)
	user_router.GET("/all-warbles", uc.FindUserWarbles)
	user_router.DELETE("/delete-warble", uc.DeleteWarble)

	r.Use(cors.Default())
	r.Run()
}
