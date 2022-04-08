package fstorage

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const (
// paths = "/data/work/gopath/src/gin-file/file"
)

// HandleDownloadFile 下载文件
func HandleDownloadFile(c *gin.Context) {

	fmt.Println("下载文件")
	//方法四
	// content := c.Query("content")
	// content = "hello zhz, 我是一个文件" + content
	// fmt.Println("content", content)
	// c.Writer.WriteHeader(http.StatusOK)
	// c.Header("Content-Disposition", "attachment; filename=zhz.txt ")
	// c.Header("Content-Type", "application/text/plain")
	// c.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
	// c.Writer.Write([]byte(content))

	//方法三
	// filename := "资产表规则.txt"
	// c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	// //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	// c.Writer.Header().Add("Content-Type", "application/octet-stream")
	// c.File("./file/test.txt")

	//方法一
	//写文件
	// var filename = "./file/test.txt"
	// if !checkFileIsExist(filename) {
	// 	file, err := os.Create(filename) //创建文件
	// 	if err != nil {
	// 		c.String(400, err.Error())
	// 		return
	// 	}
	// 	buf := bufio.NewWriter(file) //创建新的 Writer 对象
	// 	buf.WriteString("test")
	// 	buf.Flush()
	// 	defer file.Close()
	// }
	// //返回文件流
	// c.File(filename)

	//方法二
	files := c.Query("url")
	workDir, _ := os.Getwd()
	filePaths := filepath.Join(workDir, files)
	fmt.Println("filePath", filePaths)
	file, err := os.Open(filePaths) //Create a file
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "文件读取失败",
		})
		return
	}
	defer file.Close()
	filename := filepath.Base(filePaths)
	c.Writer.Header().Add("Content-type", "application/octet-stream")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "文件读取失败",
		})
		return
	}

}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
