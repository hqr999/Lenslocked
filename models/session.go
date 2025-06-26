package models

import "database/sql"

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
	//A FAZER Cria um token de sessão
	//A FAZER Implementar SessionService.Create
	return nil, nil

}

func (ss *SessionService) User(token string) (*User, error) {
	//A FAZER Implementar SessionService.user
	return nil, nil 
}
