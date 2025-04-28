package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"

)

func ParseFS(fs fs.FS, padrao string) (Template, error) {
	tpl, err := template.ParseFS(fs, padrao)
	if err != nil {
		return Template{}, fmt.Errorf("Parsisng template: %w", err)
	}

	return Template{tpl}, nil

}

func Must(te Template, err error) Template {
	if err != nil {
		panic(err)
	}

	return te
}

func ParseT(cam_arq string) (Template, error) {
	tpl, err := template.ParseFiles(cam_arq)

	if err != nil {
		return Template{}, fmt.Errorf("Parsisng template: %w", err)
	}

	return Template{tpl}, nil

}

type Template struct {
	html_tmpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	err := t.html_tmpl.Execute(w, "uma string")

	if err != nil {
		log.Printf("Executando o template: %v", err)
		http.Error(w, "Ocorreu um erro ao Executar um template", http.StatusInternalServerError)
		return //Para de rodar o código aqui
	}

}
