package controllers

import (
	"api/domain"
	"api/usecase"
	"bytes"
	"context"
	"io"
	"log"

	"github.com/google/go-cloud/blob"
)

type ImageController struct {
	Usecase usecase.ImageUsecase
}

func NewImageController() *ImageController {
	return &ImageController{
		Usecase: usecase.ImageUsecase{},
	}
}

func (controller *ImageController) Upload(c Context, blob *blob.Bucket, ctx context.Context) {
	fileHeader, _ := c.FormFile("file")
	image, _ := fileHeader.Open()
	defer image.Close()
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, image)
	w, err := blob.NewWriter(ctx, fileHeader.Filename, nil)
	if err != nil {
		c.JSON(400, NewError(err))
		return
	}
	if _, err := w.Write(buf.Bytes()); err != nil {
		c.JSON(400, NewError(err))
		return
	}
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	c.JSON(200, &domain.Image{})
}
