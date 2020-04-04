package main

import (
	"github.com/gin-gonic/gin"
	"newsfeeder/httpd/handler"
	"newsfeeder/platform/newsfeed"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default() //logger, reporting middleware, colorful

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.POST("/newsfeed", handler.NewsFeedPost(feed))

	r.Run()
}
