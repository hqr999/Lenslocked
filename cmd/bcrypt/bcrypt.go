package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	switch os.Args[1] {
	case "hash":
		//Hash a senha
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])

	default:
		fmt.Printf("Comando inválido: %v\n", os.Args[1])
	}
}

func hash(senha string) {
	hashBytes, erro := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if erro != nil {
		fmt.Printf("Erro ao fazer o hashing da senha: %v\n", senha)
		return
	}
	fmt.Printf("%v\n", string(hashBytes))
}

func compare(senha, hash string) {
	error := bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
	if error != nil {
		fmt.Printf("Senha Inválida: %v", senha)
		return
	}
	fmt.Println("Senha está correta")
}
