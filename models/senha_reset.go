package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/hqr999/Go-Web-Development/rand"
)

const (
	DuracaoDefault = 1 * time.Hour
)

type SenhaReset struct {
	ID     int
	UserID int

	//Token só é definido quando o SenhaReset é criado
	Token     string
	TokenHash string
	ExpiraEm  time.Time
}

type SenhaResetServico struct {
	BD *sql.DB
	//Bytes por token é usado para determinar quantos bytes
	//devem ser usados quando se fizer a geração do token de sessão
	//Se esse valor não for definido ou for menor que a
	//constante MinBytesPorToken, ele será ignorado e usaremos MinBytesPorToken.
	BytesPorToken int
	// Duracao é a quantidade de tempo em que SenhaReset é válido
	//Padrão é DuracaoDefault que é uma constante
	Duracao time.Duration
}

func (servico *SenhaResetServico) Cria(email string) (*SenhaReset, error) {
	//Verifica se temos um endereço de email válido
	//para um usuário, e pegar o ID desse usuário
	email = strings.ToLower(email)
	var userID int
	linha := servico.BD.QueryRow(`
			SELECT id FROM users WHERE email = $1;
		`, email)
	err := linha.Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	//Monta o reset do pw
	bytesPorToken := servico.BytesPorToken
	if bytesPorToken < MinBytesPorToken {
		bytesPorToken = MinBytesPorToken
	}
	token, err := rand.String(bytesPorToken)
	if err != nil {

	}
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	duracao := servico.Duracao
	if duracao == 0 {
		duracao = DuracaoDefault
	}

	senhaReset := SenhaReset{
		UserID:    userID,
		Token:     token,
		TokenHash: servico.hash(token),
		ExpiraEm:  time.Now().Add(duracao),
	}

	//Insere o reset da senha no BD
	linha = servico.BD.QueryRow(`
			INSERT INTO password_resets (user_id,token_hash,expires_at)
		VALUES($1,$2,$3) ON CONFLICT (user_id) DO 
		UPDATE 
		SET token_hash = $2, expires_at = $3 
		RETURNING id;`, senhaReset.UserID, senhaReset.TokenHash, senhaReset.ExpiraEm)
	err = linha.Scan(&senhaReset.ID)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return &senhaReset, nil
}

func (servico *SenhaResetServico) Consome(token string) (*User, error) {
	tokenHash := servico.hash(token)
	var usu User
	var senhaReset SenhaReset
	linha := servico.BD.QueryRow(`
				SELECT password_resets.id,
						password_resets.expires_at,
						users.id,
						users.email,
						users.password_hash
			FROM password_resets
				JOIN users ON users.id = password_resets.user_id
			WHERE password_resets.token_hash = $1;`, tokenHash)
	err := linha.Scan(&senhaReset.ID, &senhaReset.ExpiraEm,
		&usu.ID, &usu.Email, &usu.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("Consume: %w", err)
	}
	if time.Now().After(senhaReset.ExpiraEm) {
		return nil, fmt.Errorf("token expired: %w", err)
	}
	err = servico.delete(senhaReset.ID)
	if err != nil {
		return nil, fmt.Errorf("consume: %w", err)
	}
	return &usu, nil
}

func (servico *SenhaResetServico) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])

}

func (servico *SenhaResetServico) delete(id int) error {
	_, err := servico.BD.Exec(`
				DELETE FROM password_resets
				WHERE id = $1;`, id)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}
