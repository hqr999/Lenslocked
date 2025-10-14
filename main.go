package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/hqr999/Go-Web-Development/controllers"
	"github.com/hqr999/Go-Web-Development/migracoes"
	"github.com/hqr999/Go-Web-Development/models"
	"github.com/hqr999/Go-Web-Development/templates"
	"github.com/hqr999/Go-Web-Development/views"
	"github.com/joho/godotenv"
)

type config struct {
	PSQL models.PostgresConfig
	SMTP models.SMTPConfig
	CSRT struct {
		Key           string
		Secure        bool
		TrustedOrigin []string
	}
	Server struct {
		Address string
	}
}

func loadEnvConfig() (config, error) {
	var cfg config
	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}

	//A FAZER: Ler os valores de PSQL de uma var ENV
	cfg.PSQL = models.DefaultPostrgesConfig()

	//A FAZER: Ler os valores de SMTP de uma var ENV
	cfg.SMTP.Host = os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	cfg.SMTP.Port, err = strconv.Atoi(portStr)
	if err != nil {
		return cfg, nil
	}
	cfg.SMTP.Username = os.Getenv("SMTP_USERNAME")
	cfg.SMTP.Password = os.Getenv("SMTP_PASSWORD")

	//A FAZER: Ler os valores do servidor de uma var ENV
	cfg.Server.Address = ":3000"

	//A FAZER: Ler os valores de PSQL de uma var ENV
	cfg.CSRT.Key = "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	cfg.CSRT.Secure = false
	tmpVar := "localhost" + cfg.Server.Address
	cfg.CSRT.TrustedOrigin = []string{tmpVar}

	return cfg, nil
}

func main() {

	cfg, err := loadEnvConfig()
	if err != nil {
		panic(err)
	}

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
	userService := &models.UserService{
		Banco_Dados: db,
	}
	sessaoServico := &models.SessionService{
		DB: db,
	}

	senhaResetServ := &models.SenhaResetServico{
		BD: db,
	}

	emailServ := models.NovoServicoEmail(cfg.SMTP)
	galeriaServ := &models.GalleryService{
		BD: db,
	}

	//Inicializando o Middleware
	user_middleware := controllers.MiddlewareUsuario{
		SessionService: sessaoServico,
	}

	csrfMiddleware := csrf.Protect(
		[]byte(cfg.CSRT.Key),
		//TODO: Consertar antes de deploy
		csrf.Secure(cfg.CSRT.Secure),
		csrf.TrustedOrigins(cfg.CSRT.TrustedOrigin),
		csrf.Path("/"),
	)

	//Iniciando os nossos controladores
	usersC := controllers.Usuarios{
		UserService:          userService,
		SessionService:       sessaoServico,
		PasswordResetService: senhaResetServ,
		EmailService:         emailServ,
	}

	galleriecC := controllers.Galleries{
		GalleryService: galeriaServ,
	}

	tpl_pag_inscr := views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	tpl_pag_login := views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))
	tpl_pag_esq_senha := views.Must(views.ParseFS(templates.FS, "forgot-pw.gohtml", "tailwind.gohtml"))
	tpl_pag_check_email := views.Must(views.ParseFS(templates.FS, "check-your-email.gohtml", "tailwind.gohtml"))
	tpl_pag_reset_senha := views.Must(views.ParseFS(templates.FS, "reset-pw.gohtml", "tailwind.gohtml"))
	tpl_nova_pag_gal := views.Must(views.ParseFS(templates.FS, "galleries/new.gohtml", "tailwind.gohtml"))
	tpl_pag_edit_gal := views.Must(views.ParseFS(templates.FS, "galleries/edit.gohtml", "tailwind.gohtml"))
	tpl_pag_index_gal := views.Must(views.ParseFS(templates.FS, "galleries/index.gohtml", "tailwind.gohtml"))
	tpl_pag_mostra_gal := views.Must(views.ParseFS(templates.FS, "galleries/show.gohtml", "tailwind.gohtml"))

	usersC.Templates.New = tpl_pag_inscr
	usersC.Templates.Signin = tpl_pag_login
	usersC.Templates.ForgotPassword = tpl_pag_esq_senha
	usersC.Templates.CheckYourEmail = tpl_pag_check_email
	usersC.Templates.ResetPassword = tpl_pag_reset_senha
	galleriecC.Templates.New = tpl_nova_pag_gal
	galleriecC.Templates.Edit = tpl_pag_edit_gal
	galleriecC.Templates.Index = tpl_pag_index_gal
	galleriecC.Templates.Show = tpl_pag_mostra_gal

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
	r.Get("/forgot-pw", usersC.ForgotPassword)
	r.Post("/forgot-pw", usersC.ProcessForgotPassword)
	r.Get("/reset-pw", usersC.ResetPassword)
	r.Post("/reset-pw", usersC.ProcessResetPassword)
	r.Route("/users/me", func(r chi.Router) {
		r.Use(user_middleware.RequireUser)
		r.Get("/", usersC.UsuarioAtual)
	})

	r.Route("/galleries", func(r chi.Router) {
		r.Get("/{id}", galleriecC.Show)
		r.Get("/{id}/images/{filename}", galleriecC.Image)
		r.Group(func(r chi.Router) {
			r.Use(user_middleware.RequireUser)
			r.Get("/", galleriecC.Index)
			r.Get("/new", galleriecC.New)
			r.Post("/", galleriecC.Create)
			r.Get("/{id}/edit", galleriecC.Edit)
			r.Post("/{id}", galleriecC.Update)
			r.Post("/{id}/delete", galleriecC.Delete)
			r.Post("/{id}/images/{filename}/delete", galleriecC.DeleteImage)
			r.Post("/{id}/images", galleriecC.UploadImage)
		})
	})
	//r.Get("/users/me", usersC.UsuarioAtual)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Página não encontrada", http.StatusNotFound)
	})

	//Iniciando o Servidor
	fmt.Printf("Começando o servidor na porta %s... \n", cfg.Server.Address)
	err = http.ListenAndServe(cfg.Server.Address, r)
	if err != nil {
		panic(err)
	}
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
