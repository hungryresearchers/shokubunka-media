package controllers

import (
	"api/interfaces/controllers/serializer"
	"api/service"
	"bytes"
	"context"
	"io"
	"log"
	"os"

	"github.com/google/go-cloud/blob"
)

type ImageController struct {
}

func NewImageController() *ImageController {
	return &ImageController{}
}

func (controller *ImageController) Upload(c Context, blob *blob.Bucket, ctx context.Context) {
	fileHeader, _ := c.FormFile("file")
	filename := service.ChangeUniqueName(fileHeader.Filename)
	image, _ := fileHeader.Open()
	defer image.Close()
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, image)
	w, err := blob.NewWriter(ctx, filename, nil)
	if err != nil {
		c.JSON(400, NewError(err))
		return
	}
	if _, err := w.Write(buf.Bytes()); err != nil {
		c.JSON(400, NewError(err))
		return
	}
	url := "https://storage.googleapis.com/" + os.Getenv("BUCKET_NAME") + "/" + filename
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	c.JSON(200, serializer.Image{URL: url})
}
