package main

import (
	"filestore-server/conf"
	_ "filestore-server/conf"
	"filestore-server/ctrl"
	"filestore-server/db"
	"filestore-server/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 初始化配置文件
	err := conf.LoadConf(conf.GetConfPath())
	if err != nil {
		log.Fatalln("load config error", err)
	}

	// 取配置
	config := conf.GetConf()
	//初始化Mysql数据库
	_ = db.MysqlDial(&config.Mysql)
	defer db.Eloquent.Close()
	fmt.Println("MySQL connection is successful")

	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/query", handler.FileQueryHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)

	http.HandleFunc("/gorm/query", ctrl.Query)
	err = http.ListenAndServe(":8989", nil)
	if err != nil {
		fmt.Println("Failed to start server,err:%s", err.Error())
	}
	fmt.Println("Server is lisenting on loacalhost:8989")
}
