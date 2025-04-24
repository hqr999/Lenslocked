package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tlPath := filepath.Join("templates", "home.gohtml")
	executeHandler(w, tlPath)
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
	fmt.Println("Começando o servidor na porta :3000...")
	r.Get("/", homeHandler)
	r.Get("/contato", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Página não encontrada", http.StatusNotFound)
	})
	http.ListenAndServe(":3000", r)
}
