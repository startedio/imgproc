package encoder

import (
	"image"
	"image/png"
	"io"
)

// PNGEncoder returns an Encoder to PNG
func PNGEncoder() Encoder {
	return func(w io.Writer, img image.Image) error {
		return png.Encode(w, img)
	}
}
