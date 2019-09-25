package service

import (
	"filestore-server/db"
	"filestore-server/model"
	"time"
)

type FileService struct {
}

//根据sha1获取文件信息
func (file *FileService) GetFileBySha1(fileSha1 string) (info model.File, err error) {
	fileInfo := model.File{}
	fileInfo.FileSha1 = fileSha1
	if err = db.Eloquent.Where(&fileInfo).Find(&info).Error; err != nil {
		return
	}
	return
}

//文件上传
func (file *FileService) UploadFile(fileName string, fileSize int64, fileSha1 string, fileAddr string) (res model.File, err error) {
	fileSave := model.File{
		FileName: fileName,
		FileSize: fileSize,
		FileAddr: fileAddr,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		FileSha1: fileSha1,
	}

	if err = db.Eloquent.Create(&fileSave).Error; err != nil {
		return
	}
	return
}

func (file *FileService) FileList() (res []model.File, err error) {
	fileSave := model.File{}
	orm := db.Eloquent
	if err = orm.Debug().Where(&fileSave).Order("create_at desc").Find(&res).Error; err != nil {
		return
	}
	return
}

//统计
func (file *FileService) Count() (total int64, err error) {
	fileSave := model.File{}
	if err = db.Eloquent.Find(&fileSave).Count(&total).Error; err != nil {
		return
	}
	return
}
