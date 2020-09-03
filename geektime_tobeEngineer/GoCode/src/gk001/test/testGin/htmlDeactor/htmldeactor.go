package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//使用LoadHTMLGlob()或者LoadHTMLFiles()

func main(){
	router:=gin.Default()
	//router.Delims("{[{","}]}")
	router.LoadHTMLGlob("template/*")
	//router.LoadHTMLGlob("template/*")
	//router.LoadHTMLGlob("template/**/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{
			"title":"Main website",
		})



		//data:=map[string]interface{}{
		//	"lang":"Go语言",
		//	"tag":"<br>",
		//}
		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		//c.AsciiJSON(http.StatusOK,data)

		//c.JSON(200,gin.H{
		//	"message":"pong",
		//})
	})

	//router.GET("/posts/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK,"posts/index.tmpl",gin.H{
	//		"title":"Posts",
	//	})
	//})
	//router.GET("users/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK,"users/index.tmpl",gin.H{
	//		"title":"Users",
	//	})
	//})

	//html:=template.Must(template.ParseFiles("file1","file2"))
	//router.SetHTMLTemplate(html)


	//router.Run()//0.0.0.0:8080
	router.Run(":8080")//0.0.0.0:8080


}



