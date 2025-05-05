package controllers

import (
	"fmt"
	"net/http"
)

type Usuarios struct {
	Templates struct {
		New Template
	}
}

func (u Usuarios) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)

}

func (u Usuarios) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Email: ", r.FormValue("email"))
	fmt.Fprint(w, "\nPassword: " ,r.FormValue("password"))
}
