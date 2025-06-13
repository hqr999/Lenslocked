package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

func ParseFS(fs fs.FS, padroes ...string) (Template, error) {
	tpl := template.New(padroes[0])
	tpl = tpl.Funcs(
		template.FuncMap{
			"campo_csrf": func() template.HTML {
				return `<input type="hidden"/>`
			},
		},
	)

	tpl, err := tpl.ParseFS(fs, padroes...)
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

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	tpl, err := t.html_tmpl.Clone()

	if err != nil {
		log.Printf("Template clonado: %v", err)
		http.Error(w, "Houve um erro renderizando a página", http.StatusInternalServerError)
		return
	}

	tpl = tpl.Funcs(
		template.FuncMap{
			"campo_csrf": func() template.HTML {
				return csrf.TemplateField(r)
			},
		},
	)

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	err = tpl.Execute(w, data)
	
	if err != nil {
		log.Printf("Executando o template: %v", err)
		http.Error(w, "Ocorreu um erro ao Executar um template", http.StatusInternalServerError)
		return //Para de rodar o código aqui
	}

}
