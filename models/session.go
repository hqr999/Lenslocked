package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/hqr999/Go-Web-Development/rand"
)

// O mínimo de bytes necessários para serem usados em cada token de sessão
const (
	MinBytesPorToken = 32
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
	//Bytes por token é usado para determinar quantos bytes
	//devem ser usados quando se fizer a geração do token de sessão
	//Se esse valor não for definido ou for menor que a
	//constante MinBytesPorToken, ele será ignorado e usaremos MinBytesPorToken.
	BytesPorToken int
}

func (ss *SessionService) Create(userId int) (*Session, error) {
	bytesPorToken := ss.BytesPorToken
	if bytesPorToken < MinBytesPorToken {
		bytesPorToken = MinBytesPorToken
	}
	token, err := rand.String(bytesPorToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	sessao := Session{
		//Identificador será dado pelo banco de dados
		UserID:    userId,
		Token:     token,
		TokenHash: ss.hash(token),
	}
	//1. Tentar dar Update na sessão do usuário
	//2. Se obtermos um erro, criamos uma nova sessão
	linha := ss.DB.QueryRow(`
		INSERT INTO sessions
		VALUES ($1,$2) ON CONFLICT (user_id) DO 
		UPDATE
		SET token_hash = $2
		RETURNING id;`, sessao.UserID, sessao.TokenHash)
	err = linha.Scan(&sessao.ID)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return &sessao, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// 1. Hash o token de sessão
	tokenHash := ss.hash(token)
	// 2. Query para a sessão com o hash
	var user User

	// 3. Fazemos um INNER JOIN
	//entre as tabelas sessions e users
	linha := ss.DB.QueryRow(`
				SELECT users.id,users.email,users.password_hash
				FROM sessions 
				INNER JOIN users ON users.id = sessions.user_id
				WHERE sessions.token_hash = $1;
		`, tokenHash)
	err := linha.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}
	// 4. Retorna o usuário
	return &user, nil
}

func (ss *SessionService) Delete(token string) error {
	tokenHash := ss.hash(token)
	_, err := ss.DB.Exec(`
				DELETE FROM sessions 
				WHERE token_hash = $1
		`, tokenHash)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}
