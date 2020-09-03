package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r:=gin.Default()
	r.GET("/someJson", func(c *gin.Context) {
		data:=map[string]interface{}{
			"lang":"Go语言",
			"tag":"<br>",
		}
		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK,data)

		//c.JSON(200,gin.H{
		//	"message":"pong",
		//})
	})
	r.Run()//0.0.0.0:8080


}



