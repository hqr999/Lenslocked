package models

import (
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
)

type Gallery struct {
	ID     int
	UserID int
	Title  string
}

type GalleryService struct {
	BD *sql.DB

	//ImagesDir é usado para dizer ao GalleryService onde armazenar
	//e localizar imagens. Se não for setado, o GalleryService
	//vai por padrão usar o diretório 'images'.
	ImagesDir string
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
	err := linha.Scan(&galeria.Title, &galeria.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("query gallery by id: %w", err)
	}
	return &galeria, nil
}

func (gs *GalleryService) ByUserID(userID int) ([]Gallery, error) {
	linhas, erros := gs.BD.Query(`
				SELECT id, title FROM galleries
				WHERE user_id = $1;`, userID)

	if erros != nil {
		return nil, fmt.Errorf("query galleries by user id: %w", erros)

	}
	var galerias []Gallery
	for linhas.Next() {
		gal := Gallery{
			UserID: userID,
		}
		erros = linhas.Scan(&gal.ID, &gal.Title)
		if erros != nil {
			return nil, fmt.Errorf("query galleries by user id:%w", erros)

		}
		galerias = append(galerias, gal)
	}

	if linhas.Err() != nil {
		return nil, fmt.Errorf("query galleries by user id: %w", erros)
	}
	return galerias, nil
}

func (gs GalleryService) Update(gal *Gallery) error {
	_, err := gs.BD.Exec(`
				UPDATE galleries 
				SET title = $2
				WHERE id = $1;`, gal.ID, gal.Title)

	if err != nil {
		return fmt.Errorf("update gallery: %w", err)
	}
	return nil
}

func (gs GalleryService) Delete(id int) error {
	_, err := gs.BD.Exec(`
				DELETE FROM galleries
				WHERE id = $1;`, id)
	if err != nil {
		return fmt.Errorf("delete gallery: %w", err)
	}
	return nil

}

func (gs *GalleryService) galleryDirectory(id int) string {
	imagDir := gs.ImagesDir
	if imagDir == "" {
		imagDir = "images"
	}
	return filepath.Join(imagDir, fmt.Sprintf("gallery-%d", id))

}
