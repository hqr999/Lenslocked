package models

import (
	"database/sql"
	"fmt"
)

type Gallery struct {
	ID     int
	UserID int
	Title  string
}

type GalleryService struct {
	BD *sql.DB
}

func (gs GalleryService) Create(title string, userID int) (*Gallery, error) {
	galeria := Gallery{Title: title, UserID: userID}
	linha := gs.BD.QueryRow(`
				INSERT INTO galleries (title,user_id)
				VALUES ($1,$2) RETURNING id;`, galeria.Title, galeria.UserID)
	error := linha.Scan(&galeria.ID)
	if error != nil {
		return nil, fmt.Errorf("create gallery: %w", error)
	}

	return &galeria, nil
}
