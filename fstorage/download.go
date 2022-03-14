package fstorage

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	paths = "/data/work/gopath/src/gin-file/file"
)

// HandleDownloadFile 下载文件
func HandleDownloadFile(c *gin.Context) {

	fmt.Println("下载文件")
	content := c.Query("content")
	content = "hello zhz, 我是一个文件" + content
	fmt.Println("content", content)
	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", "attachment; filename=zhz.txt ")
	c.Header("Content-Type", "application/text/plain")
	c.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
	c.Writer.Write([]byte(content))

	// var filename = "/data/work/gopath/src/gin-file/file/华为1.svg"
	// file, err := os.Create(filename) //创建文件
	// if err != nil {
	// 	c.String(400, err.Error())
	// 	return
	// }
	// buf := bufio.NewWriter(file) //创建新的 Writer 对象
	// buf.WriteString("test")
	// buf.Flush()
	// defer file.Close()
	//返回文件流
	// c.File(filename)

	// filePath := "/data/work/gopath/src/gin-file/file/华为1.svg"
	// fmt.Println("sssssssss",filePath)
	// //打开文件
	// fileTmp, _ := os.Open(filePath)
	// defer fileTmp.Close()
	// //获取文件的名称
	// fileName := filepath.Join(path, "华为1.svg")
	// c.Header("Content-Type", "application/octet-stream")
	// //强制浏览器下载
	// c.Header("Content-Disposition", "attachment; filename="+fileName)
	// //浏览器下载或预览
	// c.Header("Content-Disposition", "inline;filename="+fileName)
	// c.Header("Content-Transfer-Encoding", "binary")
	// c.Header("Cache-Control", "no-cache")

	// c.File(fileName)

}
