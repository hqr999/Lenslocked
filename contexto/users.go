package contexto

import (
	"context"

	"github.com/hqr999/Go-Web-Development/models"
)

type chave string

const (
	chaveUser chave = "usuário"
)

// Função que armazena um usuário em um contexto
func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, chaveUser, user)

}

// Recupera usuários de um contexto
func User(ctx context.Context) *models.User {
	val := ctx.Value(chaveUser)
	user, ok := val.(*models.User)
	if !ok {
		//O cenário mais provável é que nada foi armazenado 
		//no contexto, então não possui um tipo *models.User. 
		//Também é possível que outro código nesse pacote
		//escreveu um valor inválido usando essa chave
		return nil
	} else {
		return user
	}
}
