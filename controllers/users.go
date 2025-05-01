package controllers

import (
	"net/http"

	"github.com/hqr999/Go-Web-Development/views"
)


type Usuarios struct{
		Templates struct {
			New views.Template
	}
}

func (u Usuarios) New(w http.ResponseWriter ,r *http.Request) {
		u.Templates.New.Execute(w,nil)

}
