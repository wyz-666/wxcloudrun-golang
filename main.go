package main

import (
	"fmt"
	"log"

	// "net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	r := gin.Default()

	r.GET("/", service.IndexHandler)
	r.GET("/api/count", service.GetCounterHandler)
	r.POST("/api/count", service.PostCounterHandler)
	r.GET("/api/user", service.GetUserInfo)
	// http.HandleFunc("/", service.IndexHandler)
	// http.HandleFunc("/api/count", service.CounterHandler)
	log.Fatal(r.Run(":80"))
	// log.Fatal(http.ListenAndServe(":80", nil))
}
