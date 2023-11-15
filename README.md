# imgproc

A comprehensive image processing utility library in Go, providing a range of functionalities for image encoding, decoding, transformation, and effects.

## Features
- Image encoding and decoding with support for popular formats like JPEG and PNG.
- Conversion utilities to standard image types in Go, like `*image.NRGBA`.
- Image effects, including StackBlur for efficient and high-quality blurring.

## Installation
To use `imgproc` in your Go project, install it using `go get`:
```shell
go get github.com/startedio/imgproc
```


## Usage

### Reading and Writing Images
The library provides simple functions to read images from different sources and write them back in various formats.

```go
import (
    "github.com/startedio/imgproc/reader"
    "github.com/startedio/imgproc/writer"
    "github.com/startedio/imgproc/encoder"
    "os"
)

func main() {
    // Reading an image
    inputFile, _ := os.Open("path/to/image.jpg")
    defer inputFile.Close()
    img, _ := reader.Read(inputFile)

    // Writing an image
    outputFile, _ := os.Create("path/to/output.png")
    defer outputFile.Close()
    _ = writer.Write(outputFile, img, encoder.PNGEncoder())
}
```

### Applying StackBlur

```go
import (    
    "github.com/startedio/imgproc/effect"
)

func main() {
    img := // obtain an image.Image
    blurredImg, err := effect.StackBlur(img, 10) // Radius: 10
}
```

## License
`imgproc` is licensed under the MIT License. See the LICENSE file for more information.

## Contributing
Contributions to `imgproc` are welcome! Please refer to our contributing guidelines for more details.

## Acknowledgements
* The StackBlur algorithm is an implementation based on the work by Endre Simó (2017) under the MIT License.
