package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hqr999/Go-Web-Development/models"
	"github.com/joho/godotenv"
)




func main()  {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Erro ao carregar o arquivo .env")
	}

	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("Erro converto a string da porta")
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	email_service := models.NovoServicoEmail(models.SMTPConfig{
			Host: host,
			Port: port,
			Username: username,
			Password: password,
	})
	
	err = email_service.EsqueceuSenha("jon@calhoun.io","https://lenslocked.com/reset-pw?token=abc123")
	if err != nil {
		panic(err)
	}

	fmt.Println("Email enviado")


}
