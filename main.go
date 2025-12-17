package main

import (
	"github.com/gin-gonic/gin"

	"gym/app"
	"gym/db"
)

func main() {
	client := db.Init()

	a := &app.App{
		Client: client,
		Key:    []byte("nuwb1icwbb1kw1cwnnc"),
	}

	r := gin.Default()

	r.POST("/adduser", a.UserHandler)
	r.POST("/addcoach", a.CoachHandler)
	r.POST("/addcourse", a.CourseHandler)
	r.POST("/userwallet", a.WalletHandler)
	r.POST("/select", a.MainHandler)
	r.POST("/login", a.LoginHandler)

	r.GET("/show", a.ShowHandler)
	r.GET("/me", a.ShowUser)

	r.Run(":8080")
}
