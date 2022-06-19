package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r:=gin.Default()
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"GET",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"PUT",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"POST",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"DELETE",
		})
	})
	r.Run(":8090")



}


