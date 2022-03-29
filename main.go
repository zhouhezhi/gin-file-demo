package main

import (
	"gin-file/controller"
	"gin-file/database"
	"gin-file/fstorage"
	"gin-file/login"
	"gin-file/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//变量
var _ = log.Printf

//启动
func main() {

	//初始化数据库

	InitConfig()
	DB := database.Init()
	defer DB.Close()
	router := gin.Default()
	router.LoadHTMLGlob("/data/work/gopath/src/gin-file-demo/static/*")

	// router.LoadHTMLFiles("/data/work/gopath/src/gin-file-demo/static/login.html")
	router.GET("/api/auth/login", func(c *gin.Context) {
		c.HTML(200, "login.html", "flysnow_org")
	})

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())

	router.POST("/api/auth/register", login.Register)
	router.POST("/api/auth/login", login.Login)
	router.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	//

	file := router.Group("/file")
	{
		file.POST("/upload", fstorage.HandleUploadFile)
		file.GET("/upload", func(c *gin.Context) {
			c.HTML(200, "upload.html", "flysnow_org")
		})
		file.POST("/upload_muti_file", fstorage.HandleUploadMutiFile)
		file.GET("/download", fstorage.HandleDownloadFile)
	}

	db := router.Group("/db")
	{
		db.POST("/book", database.CreatBook)
		db.GET("/books", database.FetchAllBooks)
		db.GET("/book/:id", database.UpdateBook)
		db.PUT("/book/:id", database.FetchSingleBook)
		db.DELETE("/book/:id", database.DeleteBook)
	}

	port := viper.GetString("server.port")
	if port != "" {
		panic(router.Run(":" + port))
	}
	panic(router.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
