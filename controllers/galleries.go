package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hqr999/Go-Web-Development/contexto"
	"github.com/hqr999/Go-Web-Development/errors"
	"github.com/hqr999/Go-Web-Development/models"
)

type Galleries struct {
	Templates struct {
		New  Template
		Edit Template
	}
	GalleryService *models.GalleryService
}

func (gal Galleries) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Title string
	}

	data.Title = r.FormValue("title")
	gal.Templates.New.Execute(w, r, data)

}

func (gal Galleries) Create(w http.ResponseWriter, r *http.Request) {
	var data struct {
		UserID int
		Title  string
	}
	data.UserID = contexto.User(r.Context()).ID
	data.Title = r.FormValue("title")

	gl, err := gal.GalleryService.Create(data.Title, data.UserID)
	if err != nil {
		gal.Templates.New.Execute(w, r, data, err)
		return
	}
	camEditar := fmt.Sprintf("/galleries/%d/edit", gl.ID)
	http.Redirect(w, r, camEditar, http.StatusFound)
}

func (gal Galleries) Edit(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválida", http.StatusNotFound)
		return
	}
	gallery, err := gal.GalleryService.ByID(id)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			http.Error(w, "Galeria não encontrada", http.StatusNotFound)
		}
		http.Error(w, "Alguma coisa deu errado", http.StatusInternalServerError)
	}
	usuario := contexto.User(r.Context())
	if gallery.ID != usuario.ID {
		http.Error(w, "Você não está autorizado para editar essa galeria", http.StatusForbidden)
		return
	}
	var data struct {
		ID    int
		Title string
	}
	data.ID = gallery.ID
	data.Title = gallery.Title

	gal.Templates.Edit.Execute(w,r,data)

}
