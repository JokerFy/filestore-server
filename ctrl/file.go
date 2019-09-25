package ctrl

import (
	"filestore-server/service"
	"filestore-server/util"
	"filestore-server/validator"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var fileService service.FileService
var FileValidate validator.FileValidate
var vErr bool

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

		fileAddr := "./tmp/" + head.Filename
		newFile, err := os.Create(fileAddr)

		if err != nil {
			util.RespFail(w, "Failed to create file,err:"+err.Error())
			return
		}
		defer newFile.Close()
		fileSize, err := io.Copy(newFile, file)
		if err != nil {
			util.RespFail(w, "Failed to save data into file,err:"+err.Error())
			return
		}

		newFile.Seek(0, 0)

		fileSha1 := util.FileSha1(newFile)
		_, err = fileService.UploadFile(head.Filename, fileSize, fileSha1, fileAddr)
		if err != nil {
			util.RespFail(w, "Failed to save data to mysql,err:"+err.Error())
			return
		}

		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}

}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "Upload finished!")
}

func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	var params = validator.FileQueryStruct{}
	fmt.Println()
	params.FileSha1 = r
	vErr = FileValidate.QueryValidate(w, params)
	if vErr {
		return
	}

	fileInfo, err := fileService.GetFileBySha1(params.FileSha1)
	if err != nil {
		util.RespFail(w, "文件不存在")
		return
	}
	util.RespOk(w, fileInfo, "文件查询成功")
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	fm, err := fileService.GetFileBySha1(fsha1)
	if err != nil {
		util.RespFail(w, "文件不存在")
		return
	}
	f, err := os.Open(string(fm.FileAddr))
	if err != nil {
		util.RespFail(w, "文件打开失败")
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		util.RespFail(w, "文件读取失败")
		return
	}

	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("content-disposition", "attachment;filename=\""+fm.FileName+"\"")
	w.Write(data)
}
