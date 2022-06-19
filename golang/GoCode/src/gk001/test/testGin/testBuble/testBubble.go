package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

// todo model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMysql() (err error) {
	//dsn := "root:104104aa@tcp(127.0.0.1:3306)/da1?charset=utf8mb&parseTime=True&loc=lcoal"
	dsn := "root:104104aa@tcp(127.0.0.1:3306)/bubble?charset=utf8&parseTime=True"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	//创建数据库
	//sql：CREAT DATABASE bubble
	//连接数据库
	err := initMysql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	//模型绑定
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		//待办事项

		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// 1.从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			// 2.存入数据库
			if err := DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
				//c.JSON(http.StatusOK,gin.H{
				//	"code":200,
				//	"msg":"success",
				//	"data":todo,
				//})
			}
			// 3.返回响应
		})
		//查看所有待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			if err := DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		//查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})

		//修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效id"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Debug().Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})

		//删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id,ok:=c.Params.Get("id")
			if !ok{
				c.JSON(http.StatusOK,gin.H{"error":"无效id"})
				return
			}
			if err = DB.Where("id=?",id).Delete(Todo{}).Error;err!=nil{
				c.JSON(http.StatusOK,gin.H{"error":err.Error()})
			}else {
				c.JSON(http.StatusOK,gin.H{id:"deleted"})
			}
		})
	}

	//r.Run(":9000")
	r.Run(":80")
}
