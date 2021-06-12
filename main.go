package main

import (
	"github.com/gin-gonic/gin"
	"github.com/litondev/gin-gionic-mangament-library/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/utrack/gin-csrf"
	// "os"	
	// "fmt"
	// "reflect"
	// "github.com/litondev/gin-gionic-mangament-library/migrations"
	// "github.com/litondev/gin-gionic-mangament-library/seeders"
)

func main(){
	// if len(os.Args) == 2 {
	// 	if os.Args[1] == "--migrate" {
	// 		migrations.RunMigration()
	// 	}

	// 	if os.Args[1] == "--seed" {
	// 		seeders.RunSeed()
	// 	}

	// 	return
	// }


	router := gin.Default()
	
	router.Static("/assets", "./assets")

	router.Use(sessions.Sessions("mycsrf", cookie.NewStore([]byte("secret"))))

	router.Use(csrf.Middleware(csrf.Options{
		Secret: "SECRET123",
		ErrorFunc: func(ctx *gin.Context) {
			ctx.String(419, "CSRF token mismatch")
			ctx.Abort()
		},
	}))

	// fmt.Println(reflect.ValueOf(router).Type())

	routers.Web(router)

 	router.Run()
}