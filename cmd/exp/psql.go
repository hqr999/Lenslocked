package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DataBaseName string
	SSLmode      string
}

func (confSQL PostgresConfig) Stringfy() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", confSQL.Host, confSQL.Port, confSQL.User, confSQL.Password, confSQL.DataBaseName, confSQL.SSLmode)
}

func main() {
	cf := PostgresConfig{"localhost", "5440", "hqr", "drachen_feuer", "lenslocked", "disable"}

	db, err := sql.Open("pgx", cf.Stringfy())

	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Conectado")

	// Criando tabelas no nosso Banco de Dados
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS users (
					id SERIAL PRIMARY KEY,
					name TEXT,
					email TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders (
				id SERIAL PRIMARY KEY,
				user_id INT NOT NULL,
				amount INT,
				description TEXT
		);
	`)

	if err != nil {
		panic(err)
	}
	fmt.Println("Tabelas criadas")

	// Inserido alguns dados na tabela
	//Veja como esse método é propenso a injection
	//O método com o $ evita a injeção, por isso sempre devemos usá-lo!!!
	name := "',''); DROP TABLE users; --"
	email := "bob@calhoun.io"
	
	//Isso causa injection
	/*query2 := fmt.Sprintf(`
			INSERT INTO users (name,email)
			VALUES ('%s','%s');
		`, name2, email2)
	fmt.Printf("Executando: %s\n", query2)
	_, err = db.Exec(query2)*/
	

	//Esse é o método que faz a verificação contra SQL Injection
	_, err = db.Exec(`
			INSERT INTO users (name,email)
			VALUES ($1,$2);`,name,email)
	if err != nil {
		panic(err)
	}

	fmt.Println("Usuário criado.")

}
