package models

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4"
)

// Vai abrir uma conexão SQL com o BD
// Postgres. Quem chama essa função precisa
// garantir que a conexão será fechada com db.Close()
func Open(config PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.String())
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	return db, nil
}

func DefaultPostrgesConfig() PostgresConfig {
	return PostgresConfig{"localhost", "5440", "hqr", "drachen_feuer", "lenslocked", "disable"}

}

type PostgresConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DataBaseName string
	SSLmode      string
}

func (confSQL PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", confSQL.Host, confSQL.Port, confSQL.User, confSQL.Password, confSQL.DataBaseName, confSQL.SSLmode)
}
