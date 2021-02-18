package uiresources

import (
	"bytes"
	"github.com/gazercloud/gazerui/canvas"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"image/png"
)

func ResImage(resName string) image.Image {
	/*iconBin, _ := Asset(resName)
	img, err := png.Decode(bytes.NewBuffer(iconBin))
	if err == nil {
		return img
	}

	img = blankImage

	return img*/
	return nil
}

func ImageFromBin(data []byte) image.Image {
	img, err := png.Decode(bytes.NewBuffer(data))
	if err == nil {
		return img
	}

	img = blankImage

	return img
}

func ImageFromBinAdjusted(data []byte, col color.Color) image.Image {
	img, err := png.Decode(bytes.NewBuffer(data))
	if err == nil {
		img = canvas.AdjustImageForColor(img, img.Bounds().Max.X, img.Bounds().Max.Y, col)
	} else {
		img = blankImage

	}
	//return nil
	return img
}

func ResBin(resName string) []byte {
	/*iconBin, _ := Asset(resName)
	return iconBin*/
	return nil
}

func ResImageAdjusted(resName string, col color.Color) image.Image {
	/*iconBin, _ := Asset(resName)
	img, err := png.Decode(bytes.NewBuffer(iconBin))
	if err == nil {
		img = canvas.AdjustImageForColor(img, img.Bounds().Max.X, img.Bounds().Max.Y, col)
	} else {
		img = blankImage

	}
	//return nil
	return img*/
	return nil
}

var blankImage *image.RGBA

func init() {
	blankImage = image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{42, 42}})
	for x := 0; x < 42; x++ {
		for y := 0; y < 42; y++ {
			blankImage.Set(x, y, colornames.Red)
		}
	}
}
