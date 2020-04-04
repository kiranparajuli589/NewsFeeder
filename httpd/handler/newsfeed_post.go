package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newsfeeder/platform/newsfeed"
)

type newsFeedPostRequest struct {
	Title string `json:"title"`
	Post string `json:"post"`
}

func NewsFeedPost(feed *newsfeed.ItemList) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsFeedPostRequest{}

		c.Bind(&requestBody) //bind incoming data from request to requestBody

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)
		c.Status(http.StatusNoContent)
	}
}
