package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// IndexHandler godoc
// @Summary Render the index page
// @Description Render the index page with HTML content
// @Tags home
// @Produce html
// @Success 200 {string} HTML content
// @Router / [get]
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
