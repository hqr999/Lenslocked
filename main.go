package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hqr999/Go-Web-Development/controllers"
	"github.com/hqr999/Go-Web-Development/models"
	"github.com/hqr999/Go-Web-Development/templates"
	"github.com/hqr999/Go-Web-Development/views"
)

func main() {
	r := chi.NewRouter()
	//Vamos parsear todos os nossos Templates
	//e depois iremos chamar nosso handlers

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))

	r.Get("/", controllers.StaticHandler(tpl))
	tpl2 := views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contato", controllers.StaticHandler(tpl2))

	tpl3 := views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl3))

	tpl4 := views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))

	tpl5 := views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))

	//Fazendo a conexão com o Banco de Dados
	config := models.DefaultPostrgesConfig()
	db, err := models.Open(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{
		Banco_Dados: db,
	}

	usersC := controllers.Usuarios{
		UserService: &userService}
	usersC.Templates.New = tpl4
	usersC.Templates.Signin = tpl5
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.Signin)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Página não encontrada", http.StatusNotFound)
	})

	fmt.Println("Começando o servidor na porta :3000...")
	http.ListenAndServe(":3000", r)
}
