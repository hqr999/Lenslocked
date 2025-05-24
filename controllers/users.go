package controllers

import (
	"fmt"
	"net/http"

	models "github.com/hqr999/Go-Web-Development/models"
)

type Usuarios struct {
	Templates struct {
		New Template
		Signin Template
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
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Criar(email, password)
	if err != nil {
		fmt.Println()
		http.Error(w, "Alguma coisa deu errado", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Usuário criado: %v", user)
}


func (u Usuarios) Signin(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.Signin.Execute(w, data)

}


func (u Usuarios) ProcessSignin(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
		Password string 
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user ,err := u.UserService.Autenticar(data.Email,data.Password)
	if err != nil {
			http.Error(w,"Alguma coisa deu errado",http.StatusInternalServerError)
		return 
	}
	fmt.Fprintf(w,"Usuário autenticado: %v", user)
}


