package main

import (
	"fmt"
	"os"

	"github.com/go-mail/mail/v2"
)

const (
	host = "sandbox.smtp.mailtrap.io"
	port = 587 
	username = "79b3bee788fc95" 
	password = "084f67c4f795be"
)

func main()  {
	de := "test@lenslocked.com"
	para := "jon@calhoun.io"
	assunto := "Esse e um e-mail de teste"
	texto := "Esse e o corpo do email"
	html := `<h1>Ola amigo!!</h1><p>Esse e o email</p><p>Espero que goste</p>`

	mensagem := mail.NewMessage()
	mensagem.SetHeader("To",de)
	mensagem.SetHeader("From",para)
	mensagem.SetHeader("Subject",assunto)
	mensagem.SetBody("text/plain",texto)
	mensagem.AddAlternative("text/html",html)
	mensagem.WriteTo(os.Stdout)
	
	dialer := mail.NewDialer(host,port,username,password)
	erro := dialer.DialAndSend(mensagem)
	if erro != nil {
		panic(erro)
	}
	fmt.Println("Mensagem enviada!!")


}
