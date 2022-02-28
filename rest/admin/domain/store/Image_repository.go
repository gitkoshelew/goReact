package store

import (
	"admin/domain/model"
	"log"
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
