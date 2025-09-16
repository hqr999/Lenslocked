package views

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/csrf"
	"github.com/hqr999/Go-Web-Development/contexto"
	"github.com/hqr999/Go-Web-Development/models"
)

type public interface {
	Public() string
}

func Must(te Template, err error) Template {
	if err != nil {
		panic(err)
	}

	return te
}

func ParseFS(fs fs.FS, padroes ...string) (Template, error) {
	tpl := template.New(path.Base(padroes[0]))
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
	return Template{
		html_tmpl: tpl,
	}, nil

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
	errMsgs := errMessages(erros...)
	tpl = tpl.Funcs(
		template.FuncMap{
			"campo_csrf": func() template.HTML {
				return csrf.TemplateField(r)
			},
			"usuarioAtual": func() *models.User {
				return contexto.User(r.Context())
			},
			"erros": func() []string {
					return errMsgs				
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

func errMessages(errs ...error) []string {
	var msgErrs []string
	for _, err := range errs {
		var pubErr public
		if errors.As(err, &pubErr) {
			msgErrs = append(msgErrs, pubErr.Public())
		} else {
			msgErrs = append(msgErrs, "Alguma coisa deu errado")
		}
	}
	return msgErrs

}
