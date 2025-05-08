package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main(){
		db, err := sql.Open("pgx","host=localhost port=5440 user=hqr password=drachen_feuer dbname=lenslocked sslmode=disable")
	if err != nil {
		panic(err)
	}
	
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Conectado")
}

