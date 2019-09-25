package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type H struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Rows  interface{} `json:"rows,omitempty"`
	Total interface{} `json:"total,omitempty"`
}

// 故事列表数据
/*type DataListStory struct {
	List  [] `json:"list"`
	Total int64   `json:"total"`
}

// 故事Response
type RespListStory struct {
	Rtn  int           `json:"rtn"`  //  0: 请求成功, 1:请求失败 -1: 需要登录
	Msg  string        `json:"msg"`  // 错误原因
	Data DataListStory `json:"data"` // 返回结果
}*/

/*func ThrowError(w http.ResponseWriter,err error){
	defer func() {
		if r := recover(); r != nil {
			log.Println("got error: ", r)
			RespFail(w,"失敗")
		}
	}()
	if err != nil {
		panic(errors.WithStack(err))
	}
}*/

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, nil, msg)
}

func RespOk(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, data, msg)
}

func RespOkList(w http.ResponseWriter, lists interface{}, total interface{}) {
	fmt.Println(lists)
	//分页数目,
	RespList(w, 0, lists, total)
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	//设置header为JSON（默认是text/html）
	w.Header().Set("Content-Type", "appliction/json")
	w.WriteHeader(http.StatusOK)
	//定义一个结构体
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	//将结构体转换成JSON输出
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err)
	}
	//输出
	w.Write(ret)
}

func RespList(w http.ResponseWriter, code int, data interface{}, total interface{}) {

	w.Header().Set("Content-Type", "application/json")
	//设置200状态
	w.WriteHeader(http.StatusOK)
	//输出
	//定义一个结构体
	//满足某一条件的全部记录数目
	//测试 100
	//20
	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	//将结构体转化成JSOn字符串
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	//输出
	w.Write(ret)
}
