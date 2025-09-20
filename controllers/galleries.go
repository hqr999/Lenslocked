package controllers

import (
	"fmt"
	"net/http"

	"github.com/hqr999/Go-Web-Development/contexto"
	"github.com/hqr999/Go-Web-Development/models"
)

type Galleries struct {
		 Templates struct {
			New Template
	}
	GalleryService *models.GalleryService 
}

func (gal Galleries) New(w http.ResponseWriter, r *http.Request){
		var data struct {
			Title string 
	}

	data.Title = r.FormValue("title")
	gal.Templates.New.Execute(w,r,data)

}

func(gal Galleries) Create(w http.ResponseWriter,r *http.Request){
		var data struct {
		UserID int 
		Title string 
	}
	data.UserID = contexto.User(r.Context()).ID
	data.Title = r.FormValue("title")

	gl,err := gal.GalleryService.Create(data.Title,data.UserID)
	if err != nil {
		gal.Templates.New.Execute(w,r,data,err)
		return 
	}
	camEditar := fmt.Sprintf("/galleries/%d/edit",gl.ID)
	http.Redirect(w,r,camEditar,http.StatusFound)
}
