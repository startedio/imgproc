package reader

import (
	"image"
	"io"
)

// Read reads an image from an io.Reader and returns an image.Image
func Read(reader io.Reader) (image.Image, error) {
	img, _, err := image.Decode(reader)

	if err != nil {
		return nil, err
	}

	return img, nil
}
