package handler

import (
	"encoding/json"
	"filestore-server/db"
	"filestore-server/meta"
	"filestore-server/model"
	"filestore-server/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传html页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			_, _ = io.WriteString(w, "internal server error")
			return
		}
		_, _ = io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//返回文件流及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data,err:%s\n", err.Error())
			return
		}
		defer file.Close()
		fileSave := model.File{
			FileName: head.Filename,
			FileAddr: "./tmp/" + head.Filename,
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
		}

		newFile, err := os.Create(fileSave.FileAddr)
		if err != nil {
			fmt.Printf("Failed to create file,err:%s\n", err.Error())
			return
		}
		defer newFile.Close()

		fileSave.FileSize, err = io.Copy(newFile, file)

		if err != nil {
			fmt.Printf("Failed to save data into file,err:%s\n", err.Error())
			return
		}

		newFile.Seek(0, 0)

		fileSave.FileSha1 = util.FileSha1(newFile)
		db.Eloquent.Debug().Create(&fileSave)

		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}

}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "Upload finished!")
}

func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileInfo := model.File{}
	fileInfo.FileSha1 = r.Form["filehash"][0]
	//fMeta := meta.GetFileMeta(filehash)
	fileInfos, err := fileInfo.GetFileBySha1()
	data, err := json.Marshal(fileInfos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func FileQueryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	limitCnt, _ := strconv.Atoi(r.Form.Get("limit"))
	fileMetas := meta.GetLastFileMetas(limitCnt)
	data, err := json.Marshal(fileMetas)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	fm := meta.GetFileMeta(fsha1)
	f, err := os.Open(fm.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("content-disposition", "attachment;filename=\""+fm.FileName+"\"")
	w.Write(data)
}
