package validator

import (
	"net/http"
)

type FileValidate struct {
}

type FileQueryStruct struct {
	FileSha1 string `validate:"required,string"`
	Status   string `validate:"required,string"`
	FileSize string `validate:"required"`
}

func (file *FileValidate) QueryValidate(w http.ResponseWriter, params FileQueryStruct) (res bool) {
	err := validate.Struct(params)
	res = validateStruct(w, err)
	return
}
