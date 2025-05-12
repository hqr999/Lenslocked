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
	
	//Vamos achar o usuário com esse id e selecionar quais 
	//colunas dele queremos 
	id := 4
	row := db.QueryRow(`
			SELECT name, email
			FROM users
			WHERE id=$1;`,id)
	
	var name,email string 
	err = row.Scan(&name,&email)
	
	//Checar se a nossa linha de retorno tem alguma coluna
	//Bom testar com um id que sabemos que não está 
	//em nosso banco de dados
	if err == sql.ErrNoRows {
		fmt.Println("Não tem uma linha para esse id = ",id)
	}

	if err != nil {
		panic(err)
	}

	fmt.Printf("User information: name=%s, email=%s \n",name,email)

	//Loop que insere 5 pedidos na tabela "orders"
	/*userID := 1

	for i := 1; i < 5; i++ {
		quantidade := i * 100
		descricao := fmt.Sprintf("Pedido Falso %d",i)
		_, err = db.Exec(`
				INSERT INTO orders(user_id, amount, description)
				VALUES($1,$2,$3)`,userID,quantidade,descricao)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Pedidos Falsos criados")*/

	type Pedido struct {
			Id int
			UserId int 
			Amount int 
			Description string 
	}
	
	var pedidos []Pedido
	userID := 1
	rows, err := db.Query(`
			SELECT id, amount, description FROM orders
			WHERE user_id=$1
		`,userID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next(){
		var pedido Pedido
		pedido.UserId = userID
		erro := rows.Scan(&pedido.Id,&pedido.Amount,&pedido.Description)

		if erro != nil {
			panic(erro)
		}

		pedidos = append(pedidos, pedido)
	}
	//O loop com rows.Next() pode retornar um erro
	//logo devemos fazer uma checagem 
	if rows.Err() != nil {
		panic(rows.Err())
	}

	for key, val := range pedidos {
		fmt.Printf("Pedido %d\n",key+1)
		fmt.Printf("ID = %d\n",val.Id)
		fmt.Printf("user_id = %d\n",val.UserId)
		fmt.Printf("Amount = %d\n",val.Amount)
		fmt.Printf("Description = %s\n",val.Description)
		fmt.Println("----------------")
	}

}
