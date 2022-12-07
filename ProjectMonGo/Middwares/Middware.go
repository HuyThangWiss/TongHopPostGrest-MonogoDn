package Middwares

import (
	"ProjectMonGo/Middwares/auth"
	"github.com/gin-gonic/gin"
)

func AuTh()gin.HandlerFunc  {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Cell")
		if tokenString == ""{
			context.JSON(401,gin.H{"Error":"request does not contain an access token"})
			context.Abort()
		}
		err:=auth.ValidateToken(tokenString)
		if err != nil{
			context.JSON(401,gin.H{"error":err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}


