package main

import (
	"ProJectTest/LuyenTap/Search/Config"
	"github.com/gin-gonic/gin"
)

func main() {
	r:=gin.Default()
	r.GET("/api/select", Config.Select)
	r.GET("/api/search",Config.Search)
	r.Run()
}
//	jsonStr := make(map[string]interface{})
//	var err1 = json.Unmarshal(body, &jsonStr)
//	if err1 != nil{
//		log.Fatalln(err1)
//		return
//	}
