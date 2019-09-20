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

	http.HandleFunc("/gorm/query", ctrl.Query)
	err = http.ListenAndServe(":8989", nil)
	if err != nil {
		fmt.Println("Failed to start server,err:%s", err.Error())
	}
	fmt.Println("Server is lisenting on loacalhost:8989")
}

// 封装f：为传入的函数增加defer
/*func errorHandler(f func()) func() {
	return func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("got error: ", r)
			}
		}()

		f()
	}
}*/

/*func main() {
	test()
}

type H struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Rows  interface{} `json:"rows,omitempty"`
	Total interface{} `json:"total,omitempty"`
}

func test() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("got error: ", r)
		}
	}()

	var err error
	err = a()
	check(err)

	err = b()
	check(err)

	_, err = c(1, 2)
	check(err)

	_, err = d(1, 0)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func a() error {
	return errors.New("error in a")
}

func b() error {
	return errors.New("error in b")
}

func c(x, y int) (int, error) {
	return x + y, errors.New("error in c")
}

func d(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("error in d, divided by 0")
	}
	return x / y, nil
}*/
