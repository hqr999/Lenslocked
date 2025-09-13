package views

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/hqr999/Go-Web-Development/contexto"
	"github.com/hqr999/Go-Web-Development/models"
)

func ParseFS(fs fs.FS, padroes ...string) (Template, error) {
	tpl := template.New(padroes[0])
	tpl = tpl.Funcs(
		template.FuncMap{
			"campo_csrf": func() (template.HTML, error) {
				return "", fmt.Errorf("Função campo_csrf ainda não foi implementado")
			},
			"usuarioAtual": func() (template.HTML, error) {
				return "", fmt.Errorf("Função usuarioAtual ainda não foi implementada")
			},
			"erros": func() []string {
				return nil
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

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}, erros ...error) {
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
			"usuarioAtual": func() *models.User {
				return contexto.User(r.Context())
			},
			"erros": func() []string {
				var errMessage []string
				for _, err := range erros {
					errMessage = append(errMessage, err.Error())
				}
				return errMessage
			},
		},
	)

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)

	if err != nil {
		log.Printf("Executando o template: %v", err)
		http.Error(w, "Ocorreu um erro ao Executar um template", http.StatusInternalServerError)
		return //Para de rodar o código aqui
	}
	io.Copy(w, &buf)

}
