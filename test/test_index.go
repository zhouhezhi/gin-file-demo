package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Web struct {
	cache_path string
}

func (w *Web) create_cache() {
	w.cache_path = filepath.Join("./", "file_server_cache")

	if _, err := os.Stat(w.cache_path); err != nil {
		// not exists
		if err := os.MkdirAll(w.cache_path, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}

func (w *Web) upload(c *gin.Context) {
	fmt.Println("enter upload functiion")
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dst := filepath.Join(w.cache_path, file.Filename)

	// ensure cache path exists
	if _, err := os.Stat(w.cache_path); err != nil {
		// not exists
		if err := os.MkdirAll(w.cache_path, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	// save file
	if err := c.SaveUploadedFile(file, dst); err != nil {
		log.Println(err)
		c.String(http.StatusNotAcceptable, fmt.Sprintf("'%s' upload failed!\n", file.Filename))
	} else {
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!\n", file.Filename))
	}
}

func (w *Web) index(c *gin.Context) {
	all_file, _ := get_all_file(w.cache_path)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"download_url": all_file})
}

func (w *Web) download_handler(c *gin.Context) {
	fmt.Println("enter downloader functiion")
	download_file, ok := c.GetPostForm("download_file")
	if !ok {
		c.String(http.StatusNotAcceptable, fmt.Sprintln("get download file from form failed"))
		return
	}
	if _, err := os.Stat(download_file); err != nil {
		c.String(http.StatusNotAcceptable, fmt.Sprintln("error, download_file not exists"))
		return
	}

	_, filename := filepath.Split(download_file)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(download_file)
}

func get_all_file(dir string) ([][]string, error) {
	dirs, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var all_file [][]string
	for _, v := range dirs {
		if !v.IsDir() {
			single_file := []string{v.Name(), filepath.Join(dir, v.Name())}
			all_file = append(all_file, single_file)
		}
	}
	return all_file, nil
}

func main() {
	w := Web{}
	w.create_cache()

	fmt.Println(w.cache_path)
	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", w.upload)
	router.GET("/", w.index)
	router.POST("/download_handler", w.download_handler)
	router.Run(":8888")
}
