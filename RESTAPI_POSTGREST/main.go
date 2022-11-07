package main

import (
	"ProJectTest/RESTAPI_POSTGREST/Config"
	"ProJectTest/RESTAPI_POSTGREST/ConnectDb"
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectDb.Connect()

	r:=gin.Default()

	r.GET("/select",Config.Select)
	r.Run()
}
