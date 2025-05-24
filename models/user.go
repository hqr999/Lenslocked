package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserService struct {
	Banco_Dados *sql.DB
}

func (u *UserService) Criar(email, password string) (*User, error) {
	email = strings.ToLower(email)
	hashBytes, erro := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if erro != nil {
		return nil, fmt.Errorf("Usuário criado: %w", erro)
	}
	passwordHash := string(hashBytes)

	usuario := User{Email: email, PasswordHash: passwordHash}
	col := u.Banco_Dados.QueryRow(`
			INSERT INTO users (email,password_hash)
			VALUES ($1, $2) RETURNING id`, email, passwordHash)
	erro = col.Scan(&usuario.ID)
	if erro != nil {
		return nil, fmt.Errorf("Usuário criado: %w", erro)
	}
	return &usuario, nil
}

func (u UserService) Autenticar(email, password string) (*User, error) {
	email = strings.ToLower(email)
	usuario := User{Email: email}

	coluna := u.Banco_Dados.QueryRow(`
			SELECT id, password_hash FROM users WHERE email=$1
		`, email)
	err := coluna.Scan(&usuario.ID, &usuario.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("Autenticação: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("Autenticação: %v", err)
	}
	fmt.Println("Senha está correta")

	return &usuario, nil
}
