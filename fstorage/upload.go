package fstorage

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	path = "/data/work/gopath/src/gin-file-demo/file"
)

func HandleUploadFile(c *gin.Context) {

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "文件上传失败了"})
		return
	}
	_, err = ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{

			"msg": "文件读取失败"})
		return
	}

	fmt.Println(header.Filename)
	//获取文件名
	fileName := filepath.Join(path, header.Filename)
	//保存文件到服务器本地
	//SaveUploadedFile(文件头，保存路径)
	if err := c.SaveUploadedFile(header, fileName); err != nil {
		c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "上传成功"})
}

func UploadFile(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("static/upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		//上传文件
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		fmt.Println("filename:", handler.Filename)
		log.Println(handler.Filename)
		log.Println(handler.Size)
		//打开文件
		f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		//存放文件
		io.Copy(f, file)
	}
}

// HandleUploadMutiFile 上传多个文件
func HandleUploadMutiFile(c *gin.Context) {

	// 限制上传文件大小
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 4<<20)
	// 限制放入内存的文件大小
	err := c.Request.ParseMultipartForm(4 << 20)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"msg": "文件读取失败"})
		return
	}
	formdata := c.Request.MultipartForm
	files := formdata.File["file"]
	// //循环存文件到服务器本地
	for _, file := range files {

		fileName := filepath.Join(path, file.Filename)

		c.SaveUploadedFile(file, fileName)
	}

	c.JSON(http.StatusOK, gin.H{

		"msg": "上传成功"})

}
