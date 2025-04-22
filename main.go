package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeHandler(w http.ResponseWriter, caminho_arquivo string) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	tpl, err := template.ParseFiles(caminho_arquivo)

	if err != nil {
		log.Printf("Erro ao parsear: %v", err)
		http.Error(w, "Ocorreu um erro ao parsear um template", http.StatusInternalServerError)
		return //Para de rodar o código aqui
	}
	err = tpl.Execute(w, "uma string")

	if err != nil {
		log.Printf("Executando o template: %v", err)
		http.Error(w, "Ocorreu um erro ao Executar um template", http.StatusInternalServerError)
		return //Para de rodar o código aqui
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
		tlPath := filepath.Join("templates","home.gohtml")
		executeHandler(w,tlPath)
}



func contactHandler(w http.ResponseWriter, r *http.Request) {
		tlPath := filepath.Join("templates","contact.html")
		executeHandler(w,tlPath)
}



func faqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1> FAQ </h1>
		<ul> 
			<li> Tem uma versão grátis ? 
						<b> Sim.Oferecemos gratuitamente por 30 dias </b>
			</li>

			<li> Qual o horário de suporte ? 
						<b> Estamos disponíveis 24 horas,7 dias por semana</b>
			</li>

			<li> Como eu entro em contato? 
				<b> Por esse <a href="um_email@proton.com">Email </a> </b>
			</li>
	</ul>
	`)
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
