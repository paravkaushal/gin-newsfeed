package main

import (
	"fmt"
	"gin-newsfeed/platform/newsfeed"
)

func main() {
	// fmt.Println("We good")

	// r := gin.Default()

	// r.GET("/ping", handler.PingGet())
	// r.Run()

	feed := newsfeed.New()

	fmt.Println(feed)

	feed.Add(newsfeed.Item{"Hello", "What are you upto"})

	fmt.Println(feed)
}
