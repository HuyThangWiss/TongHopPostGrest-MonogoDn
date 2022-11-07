package Config

import (
	"ProJectTest/RESTAPI_POSTGREST/ConnectDb"
	"ProJectTest/RESTAPI_POSTGREST/Information"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Req struct {
	Firstname string `form:"firstname"`
	Lastname  string `form:"lastname"`
	Gender    string `form:"gender"`
	Age       int    `form:"age"`
	Address   string `form:"address"`
	Gmail     string `form:"gmail"`
}

func Select(c *gin.Context) {

	ConnectDb.Connect()
	var arr []Information.Humans
	var req Req
	err := c.ShouldBindQuery(&req)
	if err != nil {
		fmt.Println(err)
		return
	}
	result :=ConnectDb.DB.Where(Information.Humans{
		Firstname: req.Firstname,
		Lastname: req.Lastname,
		Gender:    req.Gender,
		Age:       req.Age,
		Address:   req.Address,
		Gmail:     req.Gmail,
	}).Find(&arr)
	if result.Error!= nil {
		fmt.Println(result.Error.Error())
		return
	}
	c.JSON(http.StatusOK,gin.H{"data ":arr})
}