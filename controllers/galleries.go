package controllers

import (
	"net/http"

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
