package encoder

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

type Encoder func(io.Writer, image.Image) error
