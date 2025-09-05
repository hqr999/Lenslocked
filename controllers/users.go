package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/hqr999/Go-Web-Development/contexto"
	models "github.com/hqr999/Go-Web-Development/models"
)

type Usuarios struct {
	Templates struct {
		New            Template
		Signin         Template
		ForgotPassword Template
		CheckYourEmail Template
		ResetPassword  Template
	}
	UserService          *models.UserService
	SessionService       *models.SessionService
	PasswordResetService *models.SenhaResetServico
	EmailService         *models.EmailServico
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

func (u Usuarios) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.ForgotPassword.Execute(w, r, data)

}

func (u Usuarios) ProcessForgotPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	pwReset, err := u.PasswordResetService.Cria(data.Email)
	if err != nil {
		//A FAZER: Lidar com outros casos no futuro.Por exemplo, se um
		//usuário não existe com aquele endereço de e-mail
		fmt.Println(err)
		http.Error(w, "Alguma coisa deu errado", http.StatusInternalServerError)
		return
	}

	val := url.Values{
		"token": {pwReset.Token},
	}

	resetURL := "https://localhost:3000/reset-pw?" + val.Encode()
	err = u.EmailService.EsqueceuSenha(data.Email, resetURL)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Alguma coisa deu errado", http.StatusInternalServerError)
		return
	}
	//Não renderize o token de reset aqui!!
	//Precisamos que o usuário confirme que tem acesso à
	//conta de email para verificar sua identidade
	u.Templates.CheckYourEmail.Execute(w, r, data)
}

func (u Usuarios) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Token string
	}
	data.Token = r.FormValue("token")
	u.Templates.ResetPassword.Execute(w, r, data)

}

func (u Usuarios) ProcessResetPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Token    string
		Password string
	}
	data.Token = r.FormValue("token")
	data.Password = r.FormValue("password")

	user, err := u.PasswordResetService.Consome(data.Token)
	if err != nil {
		fmt.Println(err)
		//A FAZER:Distinguir entre tipos de erros
		http.Error(w, "Alguma coisa deu errado", http.StatusInternalServerError)
		return
	}

	//A FAZER: Atualizar a senha do usuário

	//O usuário agora pode entrar porque sua senha foi atualizada
	//Qualquer erros daqui em diante devem redirecionar o usuário a página de entrar(sign in)
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
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
