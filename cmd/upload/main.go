package main

import (
	"context"
	server "filestore-server/server/upload"
	"filestore-server/store/factory"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	store, err := factory.New("mysql")
	if err != nil {
		panic(err)
	}

	srv := server.NewUploadServer(":8080", store)

	errChan, err := srv.ListenAndServer()
	if err != nil {
		fmt.Println("web server start failed:", err)
		return
	}
	fmt.Println("web server start ok")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-errChan:
		fmt.Println("web server run failed:", err)
		return
	case <-c:
		fmt.Println("filestore-server program is exiting...")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Shutdown(ctx)
	}

	if err != nil {
		fmt.Println("filestore-server program exit error:", err)
		return
	}
	fmt.Println("filestore-server program exit ok")
}
