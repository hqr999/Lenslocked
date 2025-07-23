package main

import (
	"context"
	"fmt"
)

func main() {
		ctx := context.Background()
		ctx = context.WithValue(ctx,"cor favorita","azul")
		val := ctx.Value("cor favorita")
		fmt.Println(val)
}
