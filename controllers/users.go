package controllers

import (
	"net/http"
)


type Usuarios struct{
		Templates struct {
			New Template
	}
}

func (u Usuarios) New(w http.ResponseWriter ,r *http.Request) {
		u.Templates.New.Execute(w,nil)

}
