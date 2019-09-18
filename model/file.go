package model

import (
	"filestore-server/db"
	"time"
)

func (File) TableName() string {
	return "file"
}

type File struct {
	Id       int64     `gorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	FileSha1 string    `gorm:"char(40)" form:"file_sha1" json:"file_sha1"`
	FileName string    `gorm:"varchar(256)" form:"file_name" json:"file_name"`
	FileSize int64     `gorm:"bigint(20)" form:"file_size" json:"file_size"`
	FileAddr string    `gorm:"varchar(1024)" form:"file_addr" json:"file_addr"`
	CreateAt time.Time `gorm:"datetime" form:"createat" json:"createat"`
	UpdateAt time.Time `gorm:"datetime" form:"updateat" json:"updateat"`
	Status   int       `gorm:"int(11)" form:"status" json:"status"`
	Ext1     int       `gorm:"int(11)" form:"ext1" json:"ext1"`
	Ext2     int       `gorm:"int(11)" form:"ext2" json:"ext2"`
}

func (file *File) GetFileBySha1() (info File, err error) {
	if err = db.Eloquent.Where(&file).First(&info).Error; err != nil {
		return
	}
	return
}
