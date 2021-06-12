package action

import (
	"github.com/gin-gonic/gin"	
)

func Signin(ctx *gin.Context){
	ctx.String(200, "CSRF token is valid")
}

func Signup(ctx *gin.Context){
	ctx.String(200, "CSRF token is valid")
}