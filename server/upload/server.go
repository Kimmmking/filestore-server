package server

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type UploadServer struct {
	srv *http.Server
}

func NewUploadServer(addr string) *UploadServer {
	srv := &UploadServer{
		srv: &http.Server{
			Addr: addr,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/file/upload", srv.uploadHandler)

	return srv
}

func (us *UploadServer) uploadHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		// 返回上传html页面
		data, err := ioutil.ReadFile("./web/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}
		io.WriteString(w, string(data))
		// 另一种返回方式:
		// 动态文件使用http.HandleFunc设置，静态文件使用到http.FileServer设置(见main.go)
		// 所以直接redirect到http.FileServer所配置的url
		// http.Redirect(w, r, "/static/view/index.html",  http.StatusFound)
	} else if req.Method == "POST" {

	}
}
