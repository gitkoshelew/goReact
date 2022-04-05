package store

import (
	"admin/domain/model"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// ImageRepository ...
type ImageRepository struct {
	Store *Store
}

// Create image and save it to DB
func (r *ImageRepository) Create(i *model.Image) (*int, error) {
	if err := r.Store.Db.QueryRow(
		`INSERT INTO images 
		(type, ownerId) 
		VALUES ($1, $2) RETURNING id`,
		string(i.Type),
		i.OwnerID,
	).Scan(&i.ImageID); err != nil {
		r.Store.Logger.Errorf("Error occured while creating image. Err msg: %v.", err)
		return nil, err
	}
	return &i.ImageID, nil
}

// GetAll returns all images
func (r *ImageRepository) GetAll() (*[]model.Image, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM images")
	if err != nil {
		r.Store.Logger.Errorf("Error occured while getting all images. Err msg: %v.", err)
		return nil, err
	}
	images := []model.Image{}

	for rows.Next() {
		image := model.Image{}
		err := rows.Scan(
			&image.ImageID,
			&image.Type,
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
func (r *ImageRepository) FindByID(id int) (*model.ImageDTO, error) {
	image := &model.ImageDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM images WHERE id = $1",
		id).Scan(
		&image.ImageID,
		&image.Type,
		&image.OwnerID,
	); err != nil {
		r.Store.Logger.Errorf("Error occured while seacrhing image by id. Err msg: %v.", err)
		return nil, err
	}
	return image, nil
}

// Delete image from DB and local store by ID
func (r *ImageRepository) Delete(id int) error {
	var imageType string
	var ownerID int
	if err := r.Store.Db.QueryRow("DELETE FROM images WHERE id = $1 RETURNING type, ownerId", id).Scan(
		&imageType,
		&ownerID,
	); err != nil {
		r.Store.Logger.Errorf("Error occured while deleting image by id. Err msg: %v.", err)
		return err
	}

	err := os.RemoveAll(fmt.Sprintf("images/%s/id-%d", imageType, ownerID))
	if err != nil {
		r.Store.Logger.Errorf("Error occured while deleting image in local store. Err msg: %v.", err)
		return err
	}
	return nil
}

// Update image from DB
func (r *ImageRepository) Update(i *model.Image) error {

	result, err := r.Store.Db.Exec(
		`UPDATE images SET 
		type = $1 ,
		ownerId = $2 
		WHERE id = $3`,
		string(i.Type),
		i.OwnerID,
		i.ImageID,
	)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating image. Err msg: %v.", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Store.Logger.Errorf("Error occured while updating image. Err msg:%v.", err)
		return err
	}

	if rowsAffected < 1 {
		r.Store.Logger.Errorf("Error occured while updating image. Err msg:%v.", ErrNoRowsAffected)
		return ErrNoRowsAffected
	}

	r.Store.Logger.Infof("Image with id %d was updated", i.ImageID)
	return nil
}

// SaveImage image to local store and to DB
func (r *ImageRepository) SaveImage(imageDTO *model.ImageDTO, image *image.Image) (*int, error) {

	if err := r.Store.Db.QueryRow(
		`INSERT INTO images 
		(type, ownerId) 
		VALUES ($1, $2) RETURNING id`,
		imageDTO.Type,
		imageDTO.OwnerID,
	).Scan(&imageDTO.ImageID); err != nil {
		r.Store.Logger.Errorf("Error occured while saving image data in db. Err msg: %v.", err)
		return nil, err
	}

	path := fmt.Sprintf("%s/id-%d", imageDTO.Type, imageDTO.OwnerID)

	if err := os.MkdirAll(fmt.Sprintf("images/%s", path), os.ModePerm); err != nil {
		r.Store.Logger.Errorf("Error occured while creating directory for file. Err msg: %v.", err)
		return nil, err
	}

	imagesMap, err := r.ResizeImage(image)
	if err != nil {
		return nil, err
	}

	for format, img := range imagesMap {
		file, err := os.Create(fmt.Sprintf("images/%s/image-%s.jpg", path, string(format)))
		if err != nil {
			r.Store.Logger.Errorf("Error occured while creating file. Err msg: %v.", err)
			return nil, err
		}

		if err = jpeg.Encode(file, img, nil); err != nil {
			r.Store.Logger.Errorf("Error occured while encoding jpeg. Err msg: %v.", err)
			return nil, err
		}
		defer file.Close()
	}

	return &imageDTO.ImageID, nil
}

// ResizeImage ...
func (r *ImageRepository) ResizeImage(original *image.Image) (map[model.ImageFormat]image.Image, error) {
	if original == nil {
		r.Store.Logger.Errorf("Original image is nil. Err msg: %v.", ErrNilPointer)
		return nil, ErrNilPointer
	}
	images := make(map[model.ImageFormat]image.Image)
	images[model.FormatOriginal] = *original
	images[model.FormatQVGA] = resize.Resize(320, 240, *original, resize.Lanczos3)
	images[model.FormatVGA] = resize.Resize(640, 480, *original, resize.Lanczos3)
	images[model.FormatHD720p] = resize.Resize(1280, 720, *original, resize.Lanczos3)

	return images, nil
}

// GetImageFromLocalStore ...
func (r *ImageRepository) GetImageFromLocalStore(imageDTO *model.ImageDTO) (*image.Image, error) {
	path := fmt.Sprintf("images/%s/id-%d/image-%s.jpg", imageDTO.Type, imageDTO.OwnerID, imageDTO.Format)

	file, err := os.Open(path)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while opening image file. Err msg: %v.", err)
		return nil, err
	}

	image, err := jpeg.Decode(file)
	if err != nil {
		r.Store.Logger.Errorf("Error occured while decoding file to jpeg. Err msg: %v.", err)
		return nil, err
	}

	return &image, nil
}

// ModelFromDTO ...
func (r *ImageRepository) ModelFromDTO(dto *model.ImageDTO) (*model.Image, error) {

	return &model.Image{
		ImageID: dto.ImageID,
		Type:    model.ImageType(dto.Type),
		OwnerID: dto.OwnerID,
	}, nil
}
