package encoder

import (
	"image"
	"image/jpeg"
	"io"
)

// JPEGEncoder returns an Encoder to JPG with set output quality
func JPEGEncoder(quality int) Encoder {
	return func(w io.Writer, img image.Image) error {
		return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
	}
}
