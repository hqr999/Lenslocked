package main

import (
	"context"
	"fmt"

	"github.com/hqr999/Go-Web-Development/contexto"
	"github.com/hqr999/Go-Web-Development/models"
)

type ctxChave string

const (
	corFavoritaChave ctxChave = "cor favorita"
)

func main() {
	ctx := context.Background()

	usuario := models.User{
		Email: "jon@calhoun.io",
	}

	ctx = contexto.WithUser(ctx, &usuario)

	usuarioRecuperado := contexto.User(ctx)
	fmt.Println(usuarioRecuperado.Email)
}
