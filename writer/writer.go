package writer

import (
	"github.com/startedio/imgproc/encoder"
	"image"
	"io"
)

// Write writes an image to an io.Writer using the provided ImageEncoder
func Write(writer io.Writer, img image.Image, encoder encoder.Encoder) error {
	return encoder(writer, img)
}
