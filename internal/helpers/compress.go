package helpers

import (
	"bytes"
	"image"
	"image/jpeg"
	_ "image/png"
	"mime/multipart"
)

func CompressImage(file multipart.File, contentType string) (*bytes.Buffer, string, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)

	switch contentType {
	case "image/png":
		// convert PNG to JPEG to save size
		err = jpeg.Encode(buf, img, &jpeg.Options{Quality: 50})
		contentType = "image/jpeg"
	case "image/jpeg", "image/jpg":
		err = jpeg.Encode(buf, img, &jpeg.Options{Quality: 50})
	default:
		// other types (e.g. GIF) - skip
		return nil, "", err
	}

	if err != nil {
		return nil, "", err
	}

	return buf, contentType, nil
}
