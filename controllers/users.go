package controllers

import (
	"fmt"
	"net/http"

	models "github.com/hqr999/Go-Web-Development/models"
)

type Usuarios struct {
	Templates struct {
		New Template
	}
	UserService *models.UserService
}

func (u Usuarios) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)

}

func (u Usuarios) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Email: ", r.FormValue("email"))
	fmt.Fprint(w, "\nPassword: " ,r.FormValue("password"))
}
