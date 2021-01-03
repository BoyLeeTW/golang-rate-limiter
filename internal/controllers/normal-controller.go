package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NormalController struct {
}

func NewNormalController() *NormalController {
	return &NormalController{}
}

func (r *NormalController) Get(c *gin.Context) {
	if value, exist := c.Get("ip_request_count"); exist {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"body": value,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"Message:": "IntervalServerError",
	})

	return
}
