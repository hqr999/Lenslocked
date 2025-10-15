package models

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
)

var (
	ErrEmailTaken = errors.New("models: endereço de email já em uso")
	ErrNotFound   = errors.New("models: Nenhum recurso encontrado com a informação")
)

type FileError struct {
	Issue string
}

func (fe FileError) Error() string {
	return fmt.Sprintf("Arquivo inválido: %v", fe.Issue)
}

func checkContentType(r io.ReadSeeker, tiposPermitidos []string) error {
	testBytes := make([]byte, 512)
	_, err := r.Read(testBytes)
	if err != nil {
		return fmt.Errorf("Checando tipo do conteúdo: %w", err)
	}
	_, err = r.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("Checando tipo do conteúdo: %w", err)
	}
	tipoConteudo := http.DetectContentType(testBytes)
	for _, t := range tiposPermitidos {
		if tipoConteudo == t {
			return nil
		}
	}
	return FileError{Issue: fmt.Sprintf("Tipo de conteúdo inválido: %v", tipoConteudo)}
}



func checkExtensions(filename string, extenPermitidas []string) error{
		if checkExtension(filename,extenPermitidas){
				return nil 
		}
		return FileError{
			Issue: fmt.Sprintf("Extensão inválida: %v",filepath.Ext(filename)),
		}
}
