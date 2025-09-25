package main

import (
	"fmt"

	"github.com/hqr999/Go-Web-Development/models"
)


func main() {
	gs := models.GalleryService{}
	fmt.Println(gs.Images(2))
	
}
