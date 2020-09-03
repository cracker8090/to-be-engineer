package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r :=gin.Default()
	type msg struct {
		Name string `json:"name"`
		Message string
		Age int
	}
	r.GET("/json", func(c *gin.Context) {
		data :=msg{
			"小王子",
			"hello golang",
			20,
		}
		c.JSON(http.StatusOK,data)
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"message":"hello world"})
	})

	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// /JSONP?callback=x
		// 将输出：x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username","小王子")
		address := c.Query("address")
		c.JSON(http.StatusOK,gin.H{
			"message":"ok",
			"username":username,
			"address":address,
		})
	})

	r.Run(":9000")

}
