package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//Not Escaped
	//bio := `<script>alert("Haha você foi h4x0r3d!");</script>`
	
	//Escaped HTML Characters 
	bio := `&lt;script&gt;alert(&quot;Hi!&quot;);&lt;/script&gt;`

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprint(w, "<h1> Bem Vindo ao meu Site!!!</h1><p>Bio"+bio+"</p>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprint(w, `<h1>Página de Contato</h1>
		<p> Entre em contato comigo por esse e-mail: <a href="mailto:henriquereuter46@gmail.com">henriquereuter46@gmail.com</a></p>`)
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
