package models

import (
	"database/sql"
	"fmt"
	"io/fs"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
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

// Função que roda as migrações com o Goose
func Migrando(db *sql.DB, diretorio string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("Problemas com migração: %w", err)
	}
	err = goose.Up(db, diretorio)
	if err != nil {
		return fmt.Errorf("Problemas com migração: %w", err)
	}
	return nil
}

func Migrando_FS(db *sql.DB, migracoesFS fs.FS, diretorio string) error {
	if diretorio == "" {
			diretorio = "."
	}
	goose.SetBaseFS(migracoesFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()
	return Migrando(db, diretorio)
}
