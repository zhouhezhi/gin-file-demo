package main

import (
	"gin-file/database"
	"gin-file/fstorage"
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

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	file := router.Group("/file")
	{
		file.POST("/upload", fstorage.HandleUploadFile)
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

	router.Run(":8888")
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
