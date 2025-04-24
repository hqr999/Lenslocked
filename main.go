package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/hqr999/Go-Web-Development/controllers"
	"github.com/hqr999/Go-Web-Development/views"
)

func executeHandler(w http.ResponseWriter, caminho_arquivo string) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	tpl, err := views.ParseT(caminho_arquivo)
	if err != nil {
		log.Printf("Parsing template: %w", err)
		http.Error(w, "Houve um erro ao parsear o template", 404)
	}

	tpl.Execute(w, nil)

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tlPath := filepath.Join("templates", "contact.gohtml")
	executeHandler(w, tlPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tlPath := filepath.Join("templates", "faq.gohtml")
	executeHandler(w, tlPath)
}

func main() {
	r := chi.NewRouter()
	//Vamos parsear todos os nossos Templates
	//e depois iremos chamar nosso handlers
	tpl, err := views.ParseT(filepath.Join("templates", "home.gohtml"))

	if err != nil {
		panic(err)
	}

	r.Get("/", controllers.StaticHandler(tpl))
	tpl2, err_contact := views.ParseT(filepath.Join("templates", "contact.gohtml"))
	if err_contact != nil {
		panic(err_contact)
	}
	r.Get("/contato", controllers.StaticHandler(tpl2))

	tpl3, err_faq := views.ParseT(filepath.Join("templates", "faq.gohtml"))
	if err_faq != nil {
		panic(err_faq)
	}
	r.Get("/faq", controllers.StaticHandler(tpl3))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Página não encontrada", http.StatusNotFound)
	})

	fmt.Println("Começando o servidor na porta :3000...")
	http.ListenAndServe(":3000", r)
}
