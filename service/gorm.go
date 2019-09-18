package service

import (
	"filestore-server/db"
	"filestore-server/model"
	"fmt"
)

func Query(user model.User) {
	db.Eloquent.First(&user)
	fmt.Println(user)
}
