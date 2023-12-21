package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener-app/api"
)

func ShortenURLHandler(c *gin.Context) {
	userURL := c.PostForm("url")

	shortenedURL, err := api.ShortenURL(userURL)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "result.tmpl", gin.H{"url": shortenedURL})
}
