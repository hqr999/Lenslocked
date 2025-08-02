package main

import (
	"os"

	"github.com/wneessen/go-mail"
)

func main()  {
	de := "test@lenslocked.com"
	para := "jon@calhoun.io"
	assunto := "Esse e um e-mail de teste"
	texto := "Esse e o corpo do email"
	html := `<h1>Ola amigo!!</h1><p>Esse e o email</p><p>Espero que goste</p>`

	mensagem := mail.NewMsg()
	mensagem.From(de)
	mensagem.To(para)
	mensagem.Subject(assunto)
	mensagem.SetBodyString("text/plain",texto)
	mensagem.AddAlternativeString("text/html",html)
	mensagem.WriteTo(os.Stdout)
}
