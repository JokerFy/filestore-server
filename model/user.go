package model

import "time"

const (
	SEX_WOMEN  = "W"
	SEX_MEN    = "M"
	SEX_UNKNOW = "U"
)

func (User) TableName() string {
	return "user"
}

type User struct {
	//用户ID
	Id       int64  `gorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Mobile   string `gorm:"varchar(20)" form:"mobile" json:"mobile"`
	Passwd   string `gorm:"varchar(40)" form:"passwd" json:"-"` // 什么角色
	Avatar   string `gorm:"varchar(150)" form:"avatar" json:"avatar"`
	Sex      string `gorm:"varchar(2)" form:"sex" json:"sex"`            // 什么角色
	Nickname string `gorm:"varchar(20)" form:"nickname" json:"nickname"` // 什么角色
	//加盐随机字符串6
	Salt   string `gorm:"varchar(10)" form:"salt" json:"-"`    // 什么角色
	Online int    `gorm:"int(10)" form:"online" json:"online"` //是否在线
	//前端鉴权因子,
	Token    string    `gorm:"varchar(40)" form:"token" json:"token"`    // 什么角色
	Memo     string    `gorm:"varchar(140)" form:"memo" json:"memo"`     // 什么角色
	Createat time.Time `gorm:"datetime" form:"createat" json:"createat"` // 什么角色
}
