package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(templ Template) http.HandlerFunc {
	perguntas := []struct {
		Pergunta string
		Resposta template.HTML
	}{
		{
			Pergunta: "Tem uma versão grátis ?",
			Resposta: "Sim.Oferecemos gratuitamente por 30 dias.",
		},
		{
			Pergunta: "Qual o horário de suporte ?",
			Resposta: "Estamos disponíveis 24 horas,7 dias por semana.",
		},
		{
			Pergunta: "Como eu entro em contato ?",
			Resposta: `Nos envie um e-mail para - <b><a href="henriquereuter555@hotmail.com"></b>henriquereuter555@hotmail.com</a>`,
		},
		{
			Pergunta: "Onde o escritório fica localizado ?",
			Resposta: "Nós trabalhamos exclusivamente de maneira remota!!",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		templ.Execute(w, perguntas)

	}
}
