package service

import (
	"filestore-server/model"
)

func GetFileBySha1(file model.File) (info model.File, err error) {
	info, err = file.GetFileBySha1()
	if err != nil {
		return
	}
	return
}
