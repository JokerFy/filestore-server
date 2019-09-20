package main

import (
	"filestore-server/conf"
	_ "filestore-server/conf"
	"filestore-server/ctrl"
	"filestore-server/db"
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

	http.HandleFunc("/file/upload", ctrl.UploadHandler)
	http.HandleFunc("/file/upload/suc", ctrl.UploadSucHandler)
	http.HandleFunc("/file/meta", ctrl.GetFileMetaHandler)
	http.HandleFunc("/file/download", ctrl.DownloadHandler)

	err = http.ListenAndServe(":8989", nil)
	if err != nil {
		fmt.Println("Failed to start server,err:%s", err.Error())
	}
	fmt.Println("Server is lisenting on loacalhost:8989")
}
