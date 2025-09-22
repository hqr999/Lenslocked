package controllers

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hqr999/Go-Web-Development/contexto"
	"github.com/hqr999/Go-Web-Development/errors"
	"github.com/hqr999/Go-Web-Development/models"
)

type Galleries struct {
	Templates struct {
		New   Template
		Edit  Template
		Index Template
		Show  Template
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
	edit_caminho := fmt.Sprintf("/galleries/%d/edit", gl.ID)
	http.Redirect(w, r, edit_caminho, http.StatusFound)
}

func (gal Galleries) Edit(w http.ResponseWriter, r *http.Request) {
	gallery, err := gal.galleryByID(w, r, userMustOwnGallery)
	if err != nil {
		return
	}

	usuario := contexto.User(r.Context())
	if gallery.UserID != usuario.ID {
		http.Error(w, "Você não está autorizado para editar essa galeria", http.StatusForbidden)
		return
	}
	var data struct {
		ID    int
		Title string
	}
	data.ID = gallery.ID
	data.Title = gallery.Title

	gal.Templates.Edit.Execute(w, r, data)

}

func (gal Galleries) Update(w http.ResponseWriter, r *http.Request) {
	gallery, err := gal.galleryByID(w, r, userMustOwnGallery)
	if err != nil {
		return
	}
	usuario := contexto.User(r.Context())
	if gallery.UserID != usuario.ID {
		http.Error(w, "Você não está autorizado para editar essa galeria", http.StatusForbidden)
		return
	}
	gallery.Title = r.FormValue("title")
	err = gal.GalleryService.Update(gallery)
	if err != nil {
		http.Error(w, "Alguma coisa deu errado", http.StatusInternalServerError)
	}

	edit_caminho := fmt.Sprintf("/galleries/%d/edit", gallery.ID)
	http.Redirect(w, r, edit_caminho, http.StatusFound)
}

func (gal Galleries) Index(w http.ResponseWriter, r *http.Request) {
	type Gallery struct {
		ID    int
		Title string
	}

	var data struct {
		Galerias []Gallery
	}
	us := contexto.User(r.Context())
	gals, err := gal.GalleryService.ByUserID(us.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Alguma coisa deu errado", http.StatusInternalServerError)
		return
	}

	for _, gal_key := range gals {
		data.Galerias = append(data.Galerias, Gallery{gal_key.ID, gal_key.Title})

	}
	gal.Templates.Index.Execute(w, r, data)
}

func (gal Galleries) Show(w http.ResponseWriter, r *http.Request) {

	gallery, err := gal.galleryByID(w, r)
	if err != nil {
		return
	}
	var data struct {
		ID     int
		Title  string
		Images []string
	}

	data.ID = gallery.ID
	data.Title = gallery.Title
	for i := 0; i < 20; i++ {
		largura, altura := rand.IntN(500)+200, rand.IntN(500)+200
		catImageUrl := fmt.Sprintf("https://placecats.com/%d/%d", largura, altura)
		data.Images = append(data.Images, catImageUrl)

	}
	gal.Templates.Show.Execute(w, r, data)

}


func (gal Galleries) Delete(w http.ResponseWriter,r *http.Request) {
	gallery,err := gal.galleryByID(w,r,userMustOwnGallery)
	if err != nil {
			return 
	}

	err = gal.GalleryService.Delete(gallery.ID)
	if err != nil {
		http.Error(w,"Alguma coisa deu errado",http.StatusInternalServerError)
		return 
	}
	http.Redirect(w,r,"/galleries",http.StatusFound)
}

type galleryOpt func(http.ResponseWriter,*http.Request,*models.Gallery) error

func (gal Galleries) galleryByID(w http.ResponseWriter, r *http.Request,opts ...galleryOpt) (*models.Gallery, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "IP inválido", http.StatusNotFound)
		return nil, err
	}
	gallery, err := gal.GalleryService.ByID(id)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			http.Error(w, "Galeria não encontrada", http.StatusNotFound)
			return nil, err
		}
		http.Error(w, "Alguma coisa deu errado", http.StatusInternalServerError)
		return nil, err
	}

	for _, option := range opts {
			err = option(w,r,gallery)
			if err != nil {
					return nil,err 
		}
	}
	return gallery, nil

}

func userMustOwnGallery(w http.ResponseWriter, r *http.Request,gal *models.Gallery) error {
	user := contexto.User(r.Context())
	if gal.UserID != user.ID {
			http.Error(w,"Você não está autorizado a editar essa galeria",http.StatusForbidden)
			return fmt.Errorf("usuário não tem acesso a essa galeria")
	}
	return nil 
}
