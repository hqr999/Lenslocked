package models

import "github.com/go-mail/mail/v2"

const (
	RemetentePadrao = "suporte@lenslocked.com"
)

type SMTPConfig struct {
	Host string
	Port int 
	Username string 
	Password string 
}

func NovoServicoEmail(conf SMTPConfig) (*EmailServicos) {
		es := EmailServicos{
		//Precisamos preencher os campo do dialer 
		dialer: mail.NewDialer(conf.Host,conf.Port,conf.Username,conf.Password),
	}
	return &es
}

type EmailServicos struct {
	//O RemetentePadrao é usado como o remetente padrao
	//quando um não é fornecido para um email.
	//Isso também é usado em funções onde o email é
	//pré-determinado. Por exemplo quando esquece-se uma senha
	RemetentePadrao string

	//Campos que não serão exportados
	dialer *mail.Dialer
}


