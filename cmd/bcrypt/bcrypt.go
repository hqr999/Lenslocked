package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "hash":
		//Hash a senha
		hash(os.Args[2])
	case "compare":
		compara(os.Args[2],os.Args[3])

	default:
			fmt.Printf("Comando inválido: %v\n", os.Args[1])
	}
}

func hash(senha string) {
	fmt.Printf("TODO: Hash a senha %q\n", senha)
}


func compara(senha,hash string) {
	fmt.Printf("TODO: Comparar a senha %q com o hash %q\n",senha, hash)

}
