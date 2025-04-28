package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hqr999/Go-Web-Development/controllers"
	"github.com/hqr999/Go-Web-Development/templates"
	"github.com/hqr999/Go-Web-Development/views"
)

func main() {
	r := chi.NewRouter()
	//Vamos parsear todos os nossos Templates
	//e depois iremos chamar nosso handlers
	tpl := views.Must(views.ParseFS(templates.FS,"home.gohtml"))

	r.Get("/", controllers.StaticHandler(tpl))
	tpl2 := views.Must(views.ParseFS(templates.FS,"contact.gohtml"))
	r.Get("/contato", controllers.StaticHandler(tpl2))

	tpl3 := views.Must(views.ParseFS(templates.FS,"faq.gohtml"))
	r.Get("/faq", controllers.StaticHandler(tpl3))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Página não encontrada", http.StatusNotFound)
	})

	fmt.Println("Começando o servidor na porta :3000...")
	http.ListenAndServe(":3000", r)
}
