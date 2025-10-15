package models

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	//"runtime"
	"strings"
)

type Image struct {
	GalleryID int
	Path      string
	Filname   string
}

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
	dir := gs.galleryDirectory(id)
	err = os.RemoveAll(dir)
	if err != nil {
		return fmt.Errorf("delete gallery images: %w", err)
	}
	return nil

}

func (gs *GalleryService) Images(galID int) ([]Image, error) {
	//A FAZER: Implementar isso.
	globPadrao := filepath.Join(gs.galleryDirectory(galID), "*")
	allFiles, erro := filepath.Glob(globPadrao)
	if erro != nil {
		return nil, fmt.Errorf("recuperando imagens da galeria: %w", erro)
	}
	var imgs []Image
	for _, ff := range allFiles {
		if checkExtension(ff, gs.extensions()) {
			imgs = append(imgs, Image{
				Path:      ff,
				Filname:   filepath.Base(ff),
				GalleryID: galID,
			})
		}
	}
	return imgs, nil
}

func (gs *GalleryService) Image(galleryID int, filename string) (Image, error) {
	imgPath := filepath.Join(gs.galleryDirectory(galleryID), filename)
	_, err := os.Stat(imgPath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return Image{}, ErrNotFound
		}

		return Image{}, fmt.Errorf("querying uma imagem: %w", err)
	}

	return Image{
		Filname:   filename,
		GalleryID: galleryID,
		Path:      imgPath,
	}, nil
}

func (gs *GalleryService) CreateImage(galleryID int, filename string, contents io.ReadSeeker) error {
	err := checkContentType(contents, gs.imageContentTypes())

	if err != nil {
		return fmt.Errorf("Criando arquivo da imagem %v: %w", filename, err)
	}

	galleryDir := gs.galleryDirectory(galleryID)
	err = os.MkdirAll(galleryDir, 0755)
	if err != nil {
		return fmt.Errorf("Criando gallery-%d imagens no diretorio: %w", galleryID, err)
	}

	imgCam := filepath.Join(galleryDir, filename)
	destino, err := os.Create(imgCam)
	if err != nil {
		return fmt.Errorf("Criando arquivo da imagem: %w", err)
	}

	defer destino.Close()
	_, err = io.Copy(destino, contents)
	if err != nil {
		return fmt.Errorf("Copiando conteúdos da imagem: %w", err)
	}
	return nil

}

func (gs *GalleryService) DeleteImage(galleryID int, filename string) error {
	image, err := gs.Image(galleryID, filename)
	if err != nil {
		return fmt.Errorf("deletando imagem: %w", err)
	}
	err = os.Remove(image.Path)
	if err != nil {
		return fmt.Errorf("deletando imagem: %w", err)
	}
	return nil
}

func (gs *GalleryService) extensions() []string {
	return []string{".png", ".jpg", ".jpeg", ".gif"}
}

func (gs *GalleryService) imageContentTypes() []string {
	return []string{"image/png", "image/jpeg", "image/gif"}
}

func (gs *GalleryService) galleryDirectory(id int) string {
	imagDir := gs.ImagesDir
	//Another way of getting the path
	/*if imagDir == "" {
		_, file_name, _, ok := runtime.Caller(1)
		if !ok {
			panic("TODO: BETTER WAY TO HANDLE THIS")
		}
		imagDir = filepath.Join(filepath.Dir(file_name), "../images")
	}*/

	//One way of getting the path
	if imagDir == "" {
		imagDir = "images"
	}

	return filepath.Join(imagDir, fmt.Sprintf("gallery-%d", id))
}

func checkExtension(ff string, extensions []string) bool {
	for _, ext := range extensions {
		ff = strings.ToLower(ff)
		ext = strings.ToLower(ext)
		if filepath.Ext(ff) == ext {
			return true
		}
	}
	return false
}
