package validator

import (
	"filestore-server/util"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func validateStruct(w http.ResponseWriter, err error) bool {
	var vErr [10]string
	var errNews string
	var count int
	if err != nil {
		for i, err := range err.(validator.ValidationErrors) {
			errNews = err.Field() + " - " + err.Tag()
			fmt.Println(errNews)
			vErr[i] = errNews
			count = i + 1
		}
		util.RespOk(w, vErr[0:count], "ok")
		return true
	}

	return false
}

func validateVariable() {

	myEmail := "joeybloggs.gmail.com"

	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}
}
