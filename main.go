package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/hqr999/Go-Web-Development/controllers"
	"github.com/hqr999/Go-Web-Development/migracoes"
	"github.com/hqr999/Go-Web-Development/models"
	"github.com/hqr999/Go-Web-Development/templates"
	"github.com/hqr999/Go-Web-Development/views"
)

func main() {

	//Fazendo a conexão com o Banco de Dados
	config := models.DefaultPostrgesConfig()
	db, err := models.Open(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = models.Migrando_FS(db, migracoes.FS, ".")
	if err != nil {
		panic(err)
	}

	//Chamando a migração
	userService := models.UserService{
		Banco_Dados: db,
	}
	sessaoServico := models.SessionService{
		DB: db,
	}

	//Inicializando o Middleware
	user_middleware := controllers.MiddlewareUsuario{
		SessionService: &sessaoServico,
	}

	csrfChave := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMiddleware := csrf.Protect(
		[]byte(csrfChave),
		//TODO: Consertar antes de deploy
		csrf.Secure(false),
		csrf.TrustedOrigins([]string{"localhost:3000"}),
	)

	//Iniciando os nossos controladores
	usersC := controllers.Usuarios{
		UserService:    &userService,
		SessionService: &sessaoServico,
	}
	tpl_pag_inscr := views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	tpl_pag_login := views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))
	usersC.Templates.New = tpl_pag_inscr
	usersC.Templates.Signin = tpl_pag_login

	//Configurando nosso roteador e nossas rotas
	r := chi.NewRouter()
	r.Use(csrfMiddleware)
	r.Use(user_middleware.SetUsuario)
	tpl_pag_home := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl_pag_home))
	tpl_pag_contato := views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contato", controllers.StaticHandler(tpl_pag_contato))
	tpl_pag_faq := views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl_pag_faq))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.Signin)
	r.Post("/signin", usersC.ProcessSignin)
	r.Post("/signout", usersC.ProcessSignOut)
	r.Route("/users/me",func(r chi.Router) {
			r.Get("/",usersC.UsuarioAtual)
			r.Get("/hello",func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w,`
					Eu
				SoU UMa 



											StRInG V1da



				L0kA


				`)
		})
	})
	//r.Get("/users/me", usersC.UsuarioAtual)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Página não encontrada", http.StatusNotFound)
	})

	//Iniciando o Servidor
	fmt.Println("Começando o servidor na porta :3000...")
	http.ListenAndServe(":3000", r)
}

// Uncomment the TimerMiddleware func and use it above in main() to see
// it in action.
// func TimerMiddleware(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		h(w, r)
// 		fmt.Println("Request time:", time.Since(start))
// 	}
// }
