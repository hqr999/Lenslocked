package controllers

import (
	"fmt"
	"net/http"

	models "github.com/hqr999/Go-Web-Development/models"
)

type Usuarios struct {
	Templates struct {
		New    Template
		Signin Template
	}
	UserService    *models.UserService
	SessionService *models.SessionService
}

func (u Usuarios) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, r, data)

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
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		//A FAZER Deveríamos dar um warning sobre não conseguir logar
		http.Redirect(w, r, "/signin", http.StatusNotFound)
		return
	}
	cookie := http.Cookie{
		Name:     "session",
		Value:    session.Token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/signin", http.StatusFound)
	fmt.Fprintf(w, "Usuário criado: %v", user)
}

func (u Usuarios) Signin(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.Signin.Execute(w, r, data)

}

func (u Usuarios) ProcessSignin(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.UserService.Autenticar(data.Email, data.Password)
	if err != nil {
		http.Error(w, "Alguma coisa deu errado", http.StatusInternalServerError)
		return
	}
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Alguma coisa deu errado", http.StatusInternalServerError)
	}
	cookie := http.Cookie{
		Name:     "session",
		Value:    session.Token,
		Path:     "/",
		HttpOnly: true, //Isso garante que cookies não possam ser acessados através do javascript
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u Usuarios) UsuarioAtual(w http.ResponseWriter, r *http.Request) {
	token, erro := r.Cookie("session")
	if erro != nil {
		fmt.Println(erro)
		http.Redirect(w, r, "/signin", http.StatusFound)
	}
	user, err := u.SessionService.User(token.Value)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
	}
	fmt.Fprintf(w, "Usuário atual:%s\n", user.Email)
}
