package models

import (
	"database/sql"
	"errors"
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

func (gs GalleryService) ByID(id int) (*Gallery, error) {
	//A FAZER:Validação para o ID 
	galeria := Gallery{ID: id}

	linha := gs.BD.QueryRow(`
			SELECT title, user_id
			FROM galleries
			WHERE id = $1;`, galeria.ID)
	err := linha.Scan(&galeria.Title,&galeria.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows){
				return nil, ErrNotFound 
		}
		return nil,fmt.Errorf("query gallery by id: %w",err)
	}
	return &galeria,nil 
}
