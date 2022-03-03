package main

import (
	"fmt"
	"gin-newsfeed/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("We good")

	r := gin.Default()

	r.GET("/ping", handler.PingGet)
	r.Run()
}
