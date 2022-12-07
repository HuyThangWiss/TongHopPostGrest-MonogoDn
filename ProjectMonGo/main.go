package main

import (
	"ProjectMonGo/Middwares"
	"ProjectMonGo/adapters/databases"
	"ProjectMonGo/api/Controller"
	"ProjectMonGo/core/Services"
	"ProjectMonGo/core/posts"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	mongoCollection := NewMongoDBCollection()
	mongoDb := databases.NewMonGoDb(mongoCollection)
	PostRepositoryPort := posts.InitUserRepositoryPort(mongoDb)
	PostService := Services. NewUserService(PostRepositoryPort)
	PostController := Controller.NewPostServiceS(PostService)
	r:= gin.Default()
	r.POST("/CreateApi/v1/post",PostController.Create)
	r.GET("/FindAllApi/get/post",PostController.FindALl)
	r.GET("/FindApi/getId/Id/:postId",PostController.FindPostId)
	r.PATCH("/UpdateApi/Update/title/:title",PostController.UpdatePost)
	r.DELETE("/api/Delete/title/:title",PostController.DeletePost)
	r.POST("/ApiHash/Generate/Token",PostController.CreatePostToken)
	r.POST("/Get/CreateToken",PostController.GenerateToken)
	r.GET("/Information/Post/",PostController.Search)
	apiRouter := r.Group("/API",Middwares.AuTh())
	{
		apiRouter.GET("/FindAllApi/get/post",PostController.FindALl)
	}
	if err := r.Run();err != nil{
		log.Fatalf("err ",err)
		return
	}
}

func NewMongoDBCollection() *mongo.Collection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://admin:f6XPinsVTx@localhost:27017/?readPreference=primary&directConnection=true&ssl=false"))
	if err != nil {
		log.Fatalf("NewMongoDB err: %v", err)
		return nil
	}
	db := client.Database("Books")
	collection := db.Collection("Book")
	return collection
}
