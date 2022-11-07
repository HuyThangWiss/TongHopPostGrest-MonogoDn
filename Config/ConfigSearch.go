package Config

import (
	"ProJectTest/LuyenTap/Search/ConnectSearch"
	"ProJectTest/LuyenTap/Search/information"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

func Select(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	DB := ConnectSearch.ConnectDBSearch()
	postCollection := ConnectSearch.GetCollectionsearch(DB, "Humans")
	var result []information.Humans // make
	//result = make([]information.Humans, 0)
	defer cancel()
	curson2, err := postCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = curson2.All(ctx, &result); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"Data ": map[string]interface{}{"Data": result}})
}

//getCollecttion.GetCollection(DB,"Humans")

type Req struct {
	Firstname string `form:"Firstname" bson:"Firstname"`
	Lastname  string `form:"Lastname" bson:"Lastname"`
	Gender    string `form:"Gender" bson:"Gender"`
	Age       int    `form:"Age" bson:"Age"`
	Address   string `form:"Address" bson:"Address"`
	Gmail     string `form:"Gmail" bson:"Gmail"`
}

func Search(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = ConnectSearch.ConnectDBSearch()
	var postCollection = ConnectSearch.GetCollectionsearch(DB, "Humans")
	var result []information.Humans
	defer cancel()
	var req Req
	err := c.ShouldBindQuery(&req)
	if err != nil {
		fmt.Println(err)
		return
	}
	filter := bson.M{}
	if req.Firstname != "" {
		filter["Firstname"] = req.Firstname
	}
	if req.Lastname != "" {
		filter["Lastname "] = req.Lastname
	}
	//Todo check req nil
	cursor, err := postCollection.Find(ctx, filter, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusCreated, gin.H{"Data  find": map[string]interface{}{"Data": result}})
}
