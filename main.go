package main

import (
	"fmt"
	"log"

	// "net/http"
	"wxcloudrun-golang/app/handlers"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("register", handlers.Register)
		user.POST("login", handlers.Login)
	}
	r.GET("/", service.IndexHandler)
	r.GET("/api/count", service.GetCounterHandler)
	r.POST("/api/count", service.PostCounterHandler)
	r.GET("/api/user", service.GetUserInfo)
	log.Fatal(r.Run(":80"))
}
