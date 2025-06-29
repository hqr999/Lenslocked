package models

import (
	"database/sql"
	"fmt"

	"github.com/hqr999/Go-Web-Development/rand"
)

type Session struct {
	ID     int
	UserID int
	// Toke é somente setado quando cria-se uma nova sessão
	//quando buscamos por uma sessão, esse campo será deixado vazio
	//,uma vez que salvamos um token de sessão em nosso BD
	//e nós não podemos revertê-lo para um token em texto claro.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userId int) (*Session, error) {
	token, err := rand.SessionToken()
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	sessao := Session{
		//Identificador será dado pelo banco de dados
		UserID: userId,
		Token:  token,
		//A FAZER colocar o hash no token
	}
	//A FAZER salvar a sessao no BD
	return &sessao, nil

}

func (ss *SessionService) User(token string) (*User, error) {
	//A FAZER Implementar SessionService.user
	return nil, nil
}
