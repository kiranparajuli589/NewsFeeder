package main

import (
	"github.com/gin-gonic/gin"
	"newsfeeder/httpd/handler"
)

func main() {
	r := gin.Default() //logger, reporting middleware

	r.GET("/ping", handler.PingGet)

	r.Run()
}
