package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener-app/api"
)

// ShortenURLHandler godoc
// @Summary Shorten a given URL
// @Description Shorten a given URL using the API
// @Tags url
// @Accept application/x-www-form-urlencoded
// @Produce html
// @Param url formData string true "form data of the URL to be shortened"
// @Success 200 {string} HTML content with shortened URL
// @Failure 400 {string} HTML content with error
// @Router / [post]
// @Header 200 {string} Content-Type "text/html; charset=utf-8"
// @Header 400 {string} Content-Type "text/html; charset=utf-8"
func ShortenURLHandler(c *gin.Context) {
	userURL := c.PostForm("url")

	shortenedURL, err := api.ShortenURL(userURL)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "result.tmpl", gin.H{"url": shortenedURL})
}
