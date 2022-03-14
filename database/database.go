package database

import (
	"fmt"
	"gin-file/model"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	// "github.com/nacos-group/nacos-sdk-go/model"
	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 变量
var (
	Database *gorm.DB
)

type Book struct {
	ID   int    // 列名为 `id`
	Name string `gorm:"size:255"` // 列名为 `name`
	Port int    // 列名为 `port`
}

//变量
var _ = log.Printf

func Init() *gorm.DB {

	//创建一个数据库的连接
	var err error
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	db, err := gorm.Open(driverName, args)

	// db, err := gorm.Open("mysql", "root:ws-123456@tcp(192.168.146.92:3308)/file?charset=utf8")
	db.SingularTable(true)
	if err != nil {

		panic("failed to connect database")
	}
	//迁移the schema
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&model.User{})
	Database = db
	return Database
}

func GetDB() *gorm.DB {

	return Database
}

// 查询所有
func FetchAllBooks(c *gin.Context) {

	var dataList []Book
	Database.Find(&dataList)
	if len(dataList) <= 0 {

		c.JSON(http.StatusNotFound, gin.H{

			"status": -1, "result": nil, "message": "No todo found!"})
		return
	}
	c.JSON(200, gin.H{

		"status": 1, "result": dataList, "message": "Success"})
}

//查询
func FetchSingleBook(c *gin.Context) {

	var data Book
	id := c.Param("id")
	Database.First(&data, id)
	err := c.ShouldBind(&data)
	if err != nil {

		c.JSON(200, gin.H{

			"status": -1, "result": nil, "message": err.Error()})
	} else {

		if data.ID == 0 {

			c.JSON(http.StatusNotFound, gin.H{

				"status": http.StatusNotFound, "result": nil, "message": "No todo found!"})
			return
		}
		Database.Model(&data).Updates(&data)
		c.JSON(200, gin.H{

			"status": 1, "result": data, "message": "Success"})
	}

}

//更新
func UpdateBook(c *gin.Context) {

	var data Book
	id := c.Param("id")
	Database.First(&data, id)
	if data.ID == 0 {

		c.JSON(http.StatusNotFound, gin.H{

			"status": http.StatusNotFound, "result": nil, "message": "No todo found!"})
		return
	}
	Database.Model(&data).Updates(&data)
	c.JSON(200, gin.H{

		"status": 1, "result": data, "message": "Success"})
}

//创建
func CreatBook(c *gin.Context) {
	var book Book
	err := c.ShouldBind(&book)
	if err != nil {

		c.JSON(200, gin.H{

			"status": -1, "result": nil, "message": err.Error()})
	} else {

		Database.Save(&book)
		c.JSON(200, gin.H{

			"status": 1, "result": book, "message": "Success"})
	}
}

//删除
func DeleteBook(c *gin.Context) {

	var data Book
	todoID := c.Param("id")
	Database.First(&data, todoID)
	if data.ID == 0 {

		c.JSON(http.StatusNotFound, gin.H{

			"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	Database.Delete(&data)
	c.JSON(200, gin.H{

		"status": 1, "result": data, "message": "Success"})
}
