package main

import (
	"filestore-server/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		fmt.Println("Failed to start server,err:%s", err.Error())
	}
}
