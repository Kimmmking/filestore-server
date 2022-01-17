package upload

import (
	"context"
	"filestore-server/store"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type UploadServer struct {
	store store.FileMetaStore
	srv   *http.Server
}

func NewUploadServer(addr string, store store.FileMetaStore) *UploadServer {
	srv := &UploadServer{
		store: store,
		srv: &http.Server{
			Addr: addr,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "hello")
	})

	// 文件存取接口
	fileSub := router.PathPrefix("/file").Subrouter()
	fileSub.HandleFunc("/upload", srv.uploadHandler)
	fileSub.HandleFunc("/upload/suc", srv.uploadSucHandler)
	fileSub.HandleFunc("/meta", srv.getFileMetaHandler)
	fileSub.HandleFunc("/query", srv.fileQueryHandler)
	fileSub.HandleFunc("/download", srv.downloadHandler)
	fileSub.HandleFunc("/update", srv.fileUpdateHandler)
	fileSub.HandleFunc("/delete", srv.fileDeleteHandler)

	srv.srv.Handler = router
	return srv
}

func (us *UploadServer) ListenAndServer() (<-chan error, error) {
	var err error
	errChan := make(chan error)
	go func() {
		err = us.srv.ListenAndServe()
		errChan <- err
	}()

	select {
	case err = <-errChan:
		return nil, err
	case <-time.After(time.Second):
		return errChan, nil
	}
}

func (us *UploadServer) Shutdown(ctx context.Context) error {
	return us.srv.Shutdown(ctx)
}

// uploadHandler: 处理文件上传
func (us *UploadServer) uploadHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		// 返回上传html页面
		data, err := ioutil.ReadFile("../../web/view/index.html")
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

func (us *UploadServer) uploadSucHandler(w http.ResponseWriter, req *http.Request) {

}

func (us *UploadServer) getFileMetaHandler(w http.ResponseWriter, req *http.Request) {

}

func (us *UploadServer) fileQueryHandler(w http.ResponseWriter, req *http.Request) {

}

func (us *UploadServer) downloadHandler(w http.ResponseWriter, req *http.Request) {

}

func (us *UploadServer) fileUpdateHandler(w http.ResponseWriter, req *http.Request) {

}

func (us *UploadServer) fileDeleteHandler(w http.ResponseWriter, req *http.Request) {

}
