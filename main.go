package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
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
	sessaoServico := models.SessionService{
		DB: db,
	}
	usersC := controllers.Usuarios{
		UserService:    &userService,
		SessionService: &sessaoServico,
	}
	usersC.Templates.New = tpl4
	usersC.Templates.Signin = tpl5
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.Signin)
	r.Post("/signin", usersC.ProcessSignin)
	r.Post("/signout", usersC.ProcessSignOut)
	r.Get("/users/me", usersC.UsuarioAtual)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Página não encontrada", http.StatusNotFound)
	})

	fmt.Println("Começando o servidor na porta :3000...")

	csrfChave := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMiddleware := csrf.Protect(
		[]byte(csrfChave),
		//TODO: Consertar antes de deploy
		csrf.Secure(false),
		csrf.TrustedOrigins([]string{"localhost:3000"}),
	)

	http.ListenAndServe(":3000", csrfMiddleware(r))
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
