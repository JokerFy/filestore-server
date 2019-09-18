package ctrl

import (
	"filestore-server/db"
	"filestore-server/model"
	"io"
	"net/http"
)

func Query(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	db.Eloquent.First(&user)
	_, _ = io.WriteString(w, string(user.Mobile))
}
