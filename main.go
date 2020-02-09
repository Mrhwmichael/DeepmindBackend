package main

import (
	"github.com/gin-gonic/gin"
	"main.go/handlers"
	_ "main.go/handlers"
)

const (
	ip string = "192.168.0.108:4000"
)

func main(){

	router := gin.Default()

	router.GET("/test", func(context *gin.Context) {
		context.Writer.WriteString("success")
	})

	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	queryGroup := router.Group("/query")
	{
		queryGroup.GET("/all",handlers.HandleQuery)
	}

	trackGroup := router.Group("/tracking")
	{
		trackGroup.GET("/all",handlers.HandleTrack)
	}

	//需要把ip换成局域网ip地址
	router.Run(ip)
}