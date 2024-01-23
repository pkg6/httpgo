package t

import (
	"fmt"
	"golang.org/x/image/webp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

const (
	ImageExtJPG  = ".jpg"
	ImageExtJPEG = ".jpeg"
	ImageExtGIF  = ".gif"
	ImageExtPNG  = ".png"
	ImageExtWEBP = ".webp"
)

// ImageConfig
// fs, err := os.Open(imagePath)
// ext:=strings.ToLower(filepath.Ext(imagePath))
//ImageConfig(fs,ext)
func ImageConfig(r io.Reader, ext string) (config image.Config, err error) {
	switch ext {
	case ImageExtJPG, ImageExtJPEG:
		config, err = jpeg.DecodeConfig(r)
	case ImageExtGIF:
		config, err = gif.DecodeConfig(r)
	case ImageExtPNG:
		config, err = png.DecodeConfig(r)
	case ImageExtWEBP:
		config, err = webp.DecodeConfig(r)
	default:
		err = fmt.Errorf("unsupported image format %s", ext)
	}
	return
}
