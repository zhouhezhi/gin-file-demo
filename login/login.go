package login

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		//使用时间通过MD5生成token
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.html") //解析模板
		t.Execute(w, token)                       //渲染模板并发送
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
			fmt.Println()
		} else {
			//不存在token报错
			fmt.Println()
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		if len(r.Form["username"][0]) == 0 {
			fmt.Fprint(w, "用户名不能为空")
			fmt.Fprint(w, "\n password: ", r.FormValue("password"))
			return
		}

		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端

		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))

		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
		template.HTMLEscape(w, []byte(r.Form.Get("password")))
	}
}
