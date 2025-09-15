package models

import "errors"

var (
	ErrEmailTaken = errors.New("models: endereço de email já em uso")
	ErrNotFound = errors.New("models: Nenhum recurso encontrado com a informação")
)
