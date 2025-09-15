package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
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
		var pgError *pgconn.PgError
		if errors.As(erro, &pgError) {
			if pgError.Code == pgerrcode.UniqueViolation {
				return nil, ErrEmailTaken
			}
		}
		fmt.Printf("Type = %T\n", erro)
		fmt.Printf("Error = %v\n", erro)
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

	return &usuario, nil
}

func (us *UserService) UpdatePassword(userID int, pw string) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("update password: %w", err)
	}
	passwordHash := string(hashedBytes)
	_, err = us.Banco_Dados.Exec(`
			UPDATE users
			SET password_hash = $2
			WHERE id = $1;`, userID, passwordHash)
	if err != nil {
		return fmt.Errorf("update password: %w", err)
	}
	return nil

}
