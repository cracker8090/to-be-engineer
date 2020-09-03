package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)
//静态文件 样式文件 css js 图片
func main() {
	r := gin.Default()
	r.Static("/xxx","./statics")
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/posts/index.tmpl","templates/users/index.tmpl")*/
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "posts.hellolinux.xyz",
		})
	})
	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			//"title": "users.hellolinux.xyz",
			"title": "<a href='https://hellolinux.xyz'>北平的博客</a>",
		})
	})
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK,"home.html",nil)
	})
	r.Run(":9000")

}

//func main() {
//	router := gin.Default()
//	router.LoadHTMLGlob("templates/index.tmpl")
//	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
//	router.GET("/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "index.tmpl", gin.H{
//			"title": "hellolinux.xyz",
//		})
//	})
//	router.Run(":9000")
//}
