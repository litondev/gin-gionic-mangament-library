package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/multitemplate"
	userWeb "github.com/litondev/gin-gionic-mangament-library/controllers/user/web"
	userAction "github.com/litondev/gin-gionic-mangament-library/controllers/user/action"
)

func Web(router *gin.Engine){
	router.HTMLRender = func() multitemplate.Renderer {
		template := multitemplate.NewRenderer()
		template.AddFromFiles("user.signin","views/user/signin.html")
		template.AddFromFiles("user.signup","views/user/signup.html")
		template.AddFromFiles("user.crud.form","views/user/layout/default.html","views/user/crud/form.html")
		template.AddFromFiles("user.crud.index","views/user/layout/default.html","views/user/crud/index.html")
		template.AddFromFiles("user.crud.show","views/user/layout/default.html","views/user/crud/show.html")
		return template
	}()

	router.GET("/",userWeb.Signin)
	router.GET("/signin",userWeb.Signin)
	router.GET("/signup",userWeb.Signup)
	router.POST("/signin",userAction.Signin)
	router.POST("/singup",userAction.Signup)
}