package templates

import "embed"

//Isso vai permitir executar nosso servidor em 
//qualquer path, ou seja, os templates foram 
//encorporados ao binário da build do projeto
//Teste com: go build -o app .
//E execute de outra lugar no sistema


//go:embed *.gohtml **/*.gohtml
var FS embed.FS
