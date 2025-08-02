package models

import (
	"fmt"

	"github.com/go-mail/mail/v2"
)

const (
	RemetentePadrao = "suporte@lenslocked.com"
)

type Email struct {
	De      string
	Para    string
	Assunto string
	Texto   string
	HTML    string
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NovoServicoEmail(conf SMTPConfig) *EmailServico {
	es := EmailServico{
		//Precisamos preencher os campo do dialer
		dialer: mail.NewDialer(conf.Host, conf.Port, conf.Username, conf.Password),
	}
	return &es
}

type EmailServico struct {
	//O RemetentePadrao é usado como o remetente padrao
	//quando um não é fornecido para um email.
	//Isso também é usado em funções onde o email é
	//pré-determinado. Por exemplo quando esquece-se uma senha
	RemetentePadrao string

	//Campos que não serão exportados
	dialer *mail.Dialer
}

func (es *EmailServico) Mandar(email Email) error {
	mensagem := mail.NewMessage()
	mensagem.SetHeader("To", email.De)
	es.setFrom(mensagem,email)
	mensagem.SetHeader("Subject", email.Assunto)
	switch {
	case email.Texto != "" && email.HTML != "":
		mensagem.SetBody("text/plain", email.Texto)
		mensagem.AddAlternative("text/html", email.HTML)
	case email.Texto != "":
		mensagem.SetBody("text/plain", email.Texto)
	case email.HTML != "":
		mensagem.AddAlternative("text/html", email.HTML)
	}

	err := es.dialer.DialAndSend(mensagem)
	if err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

func (es *EmailServico) setFrom(msg *mail.Message,email Email) {
		var de string 
		switch  {
		case email.De != "":
			de = email.De
		case es.RemetentePadrao != "":
			de = es.RemetentePadrao
	default:
			de = RemetentePadrao
		}
	msg.SetHeader("from",de)
}
