package main

import (
	"context"
	"fmt"
)

type ctxChave string 

const(
	corFavoritaChave ctxChave = "cor favorita"
)

func main() {
		ctx := context.Background()
		//Colocamos a cor para ser azul
		ctx = context.WithValue(ctx,corFavoritaChave,"azul")

		
		//Outro pacote muda a cor 
		ctx = context.WithValue(ctx,"cor favorita","vermelho")

		val1 := ctx.Value(corFavoritaChave)
		val2 := ctx.Value("cor favorita")
		fmt.Println(val1,val2)
}

