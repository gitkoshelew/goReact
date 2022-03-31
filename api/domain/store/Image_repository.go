package store

import (
	"fmt"
	"goReact/domain/model"
	"image"
	"image/jpeg"
	"log"
	"os"
	"time"
)

// ImageRepository ...
type ImageRepository struct {
	Store *Store
}

// Create image and save it to DB
func (r *ImageRepository) Create(i *model.Image) (*model.Image, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO image",
		"(type, URL, ownerId)",
		"VALUES ($1, $2, $3) RETURNING id",
		string(i.Type),
		i.URL,
		i.OwnerID,
	).Scan(&i.ImageID); err != nil {
		log.Print(err)
		return nil, err
	}
	return i, nil
}

// GetAll returns all images
func (r *ImageRepository) GetAll() (*[]model.Image, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM image")
	if err != nil {
		log.Print(err)
	}
	images := []model.Image{}

	for rows.Next() {
		image := model.Image{}
		err := rows.Scan(
			&image.ImageID,
			&image.Type,
			&image.URL,
			&image.OwnerID,
		)
		if err != nil {
			log.Print(err)
			continue
		}
		images = append(images, image)
	}
	return &images, nil
}

//FindByID searchs and returns image by ID
func (r *ImageRepository) FindByID(id int) (*model.Image, error) {
	image := &model.Image{}
	if err := r.Store.Db.QueryRow("SELECT * FROM image WHERE id = $1",
		id).Scan(
		&image.ImageID,
		&image.Type,
		&image.URL,
		&image.OwnerID,
	); err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	return image, nil
}

// Delete image from DB by ID
func (r *ImageRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM image WHERE id = $1", id)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	log.Printf("Image deleted, rows affectet: %d", result)
	return nil
}

// Update image from DB
func (r *ImageRepository) Update(i *model.Image) error {

	result, err := r.Store.Db.Exec(
		"UPDATE image SET",
		"type = $1, URL = $2, ownerId = $3",
		"WHERE id = $4",
		string(i.Type),
		i.URL,
		i.OwnerID,
		i.ImageID,
	)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	log.Printf("Image updated, rows affectet: %d", result)
	return nil
}

// SaveImage image and save it to DB
func (r *ImageRepository) SaveImage(imageDTO *model.ImageDTO, image *image.Image) (*int, error) {
	path := fmt.Sprintf("images/%s/id-%d", imageDTO.Type, imageDTO.OwnerID)

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		r.Store.Logger.Errorf("Error occured while creating directory for file. Err msg: %v.", err)
		return nil, err
	}

	imageDTO.URL = fmt.Sprintf("%s/%s", path, time.Now().Format("2006-01-02-03-04-05"))
	imagesMap := r.ResizeImage(image)
	for format, img := range imagesMap {
		file, err := os.Create(fmt.Sprintf("%s-%s.jpg", imageDTO.URL, string(format)))
		if err != nil {
			r.Store.Logger.Errorf("Error occured while creating file. Err msg: %v.", err)
			return nil, err
		}

		if err = jpeg.Encode(file, img, nil); err != nil {
			r.Store.Logger.Errorf("Error occured while encoding jpeg. Err msg: %v.", err)
			return nil, err
		}
	}

	if err := r.Store.Db.QueryRow(
		`INSERT INTO images 
		(type, URL, ownerId) 
		VALUES ($1, $2, $3) RETURNING id`,
		imageDTO.Type,
		imageDTO.URL,
		imageDTO.OwnerID,
	).Scan(&imageDTO.ImageID); err != nil {
		r.Store.Logger.Errorf("Error occured while saving image data in db. Err msg: %v.", err)
		return nil, err
	}

	return &imageDTO.ImageID, nil
}

// ResizeImage ...
func (r *ImageRepository) ResizeImage(original *image.Image) map[model.ImageFormat]image.Image {
	images := make(map[model.ImageFormat]image.Image)
	images[model.FormatOriginal] = *original
	images[model.FormatQVGA] = resize.Resize(320, 240, *original, resize.Lanczos3)
	images[model.FormatVGA] = resize.Resize(640, 480, *original, resize.Lanczos3)
	images[model.FormatHD720p] = resize.Resize(1280, 720, *original, resize.Lanczos3)

	return images
}
