package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprint(w, "<h1> Bem Vindo ao meu Incrível site</h1>")
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

type Roteador struct{}

func (rote Roteador) ServeHTTP(w http.ResponseWriter, r *http.Request){
			switch r.URL.Path{
				case "/":
							homeHandler(w,r)
			case "/contato":
						contactHandler(w,r)
			case "/faq":
						faqHandler(w,r)
			default:
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprint(w,"Página não existe!")	
				
			}
			


}

func main() {
	fmt.Println("Ouvindo na porta :3000...")
	var roteador Roteador
	http.ListenAndServe(":3000",roteador)
}
