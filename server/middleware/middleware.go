package middleware

import (
	"mime"
	"net/http"
)

func Validating(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		contentType := req.Header.Get("Content-Type")
		mediatype, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if mediatype != "application/json" {
			http.Error(w, "invalid Content-Type", http.StatusUnsupportedMediaType)
			return
		}
		next.ServeHTTP(w, req)
	})
}
