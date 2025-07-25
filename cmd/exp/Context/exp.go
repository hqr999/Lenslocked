package main

import (
	"context"
	"fmt"
	"strings"
)

type ctxChave string

const (
	corFavoritaChave ctxChave = "cor favorita"
)

func main() {
	ctx := context.Background()
	//Colocamos a cor para ser azul
	ctx = context.WithValue(ctx, corFavoritaChave, "azul")

	val := ctx.Value(corFavoritaChave)

	intV, ok := val.(int)

	if !ok {
		fmt.Println("Não é um inteiro")
	} else {
		fmt.Println(intV + 4)
	}

	strV, ok := val.(string)
	if !ok {
		fmt.Println("Não é uma string")
	} else {
		fmt.Println(strings.HasPrefix(strV, "a"))
	}
}
