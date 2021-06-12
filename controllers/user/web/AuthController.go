package web

import (
	"github.com/gin-gonic/gin"
	"github.com/utrack/gin-csrf"
)

func Signin(ctx *gin.Context){
	ctx.HTML(200,"user.signin",gin.H{
		"csrfToken" : csrf.GetToken(ctx),
	})
}

func Signup(ctx *gin.Context){
	ctx.HTML(200,"user.signup",gin.H{
		"csrfToken" : csrf.GetToken(ctx),
	})
}