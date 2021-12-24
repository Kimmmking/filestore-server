package main

import (
	"github.com/Kimmmking/filestore-server/server"
)

func main() {
	srv := server.NewUploadServer(":8080")

}
