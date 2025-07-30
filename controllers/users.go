package controllers

import (
	"fmt"
	"net/http"

	"github.com/hqr999/Go-Web-Development/contexto"
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
	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
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
	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u Usuarios) UsuarioAtual(w http.ResponseWriter, r *http.Request) {
	contxt := r.Context()
	usuario := contexto.User(contxt)

	if usuario == nil {
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	fmt.Fprintf(w, "Usuário atual:%s\n", usuario.Email)
}

func (u Usuarios) ProcessSignOut(w http.ResponseWriter, r *http.Request) {
	token, err := readCookie(r, CookieSession)
	if err != nil {
		http.Redirect(w, r, "/signin", http.StatusFound)
	}
	err = u.SessionService.Delete(token)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Alguma coisa deu errado", http.StatusInternalServerError)
		return
	}
	deleteCookie(w, token)
	http.Redirect(w, r, "/signin", http.StatusFound)
}

type MiddlewareUsuario struct {
	SessionService *models.SessionService
}

func (us_mw MiddlewareUsuario) SetUsuario(prox http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, erro := readCookie(r, CookieSession)
		if erro != nil {
			prox.ServeHTTP(w, r)
			return
		}
		user, err := us_mw.SessionService.User(token)
		if err != nil {
			prox.ServeHTTP(w, r)
			return
		}
		contxt := r.Context()
		contxt = contexto.WithUser(contxt, user)
		r = r.WithContext(contxt)
		prox.ServeHTTP(w, r)
	})

}

func (us_mw MiddlewareUsuario) RequireUser(prox http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usuario := contexto.User(r.Context())
		if usuario == nil {
			http.Redirect(w, r, "/signin", http.StatusFound)
		}
		prox.ServeHTTP(w, r)
	})
}
