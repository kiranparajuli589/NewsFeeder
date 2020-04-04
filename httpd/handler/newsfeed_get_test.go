package handler

import (
	"net/http"
	"newsfeeder/platform/newsfeed"
	"testing"
)

func TestNewsFeedGet(t *testing.T) {
	feed := newsfeed.New()
	feed.Add(Item{})
	resp, err := http.Get("http://localhost:8080/newsfeed")
	if err != nil {
		println(err)
	}
	print(resp)
}
