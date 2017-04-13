package processor

import (
	"bytes"
	"image"
	"image/color"
	"reflect"

	"github.com/nfnt/resize"
)

type AsciiImage struct {
	Image  string `json:"image"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

const MAX_WIDTH = 240
const MAX_HEIGHT = 240
const ASCIIMAP = "MN#80Z$ZY?+=~:,. "

func processPixel(px color.Color) uint64 {
	discolored := color.GrayModel.Convert(px)
	value := reflect.ValueOf(discolored).FieldByName("Y")

	return value.Uint()
}

func charPosFromValue(value uint64) byte {
	pos := int(value * 16 / 255)

	return ASCIIMAP[pos]
}

func scaleImage(img image.Image) image.Image {
	// set lower value to 0
	// for resize to aspect preserve ratio
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y

	if w > h {
		h = 0
		w = MAX_WIDTH
	} else {
		w = 0
		h = MAX_HEIGHT
	}

	img = resize.Resize(uint(w), uint(h), img, resize.NearestNeighbor)

	return img
}

func (ai *AsciiImage) convert(img image.Image) {
	var buffer bytes.Buffer

	for i := 0; i < ai.Height; i++ {
		for j := 0; j < ai.Width; j++ {
			pixelValue := processPixel(img.At(j, i))
			buffer.WriteByte(charPosFromValue(pixelValue))
		}

		buffer.WriteByte('\n')
	}

	ai.Image = buffer.String()
}

func NewAsciiImage(src image.Image) *AsciiImage {
	bounds := src.Bounds().Max
	asciiImage := &AsciiImage{
		Width:  bounds.X,
		Height: bounds.Y,
	}
	asciiImage.convert(scaleImage(src))

	return asciiImage
}

func Convert(img image.Image) *AsciiImage {
	img = scaleImage(img)
	ai := NewAsciiImage(img)

	return ai
}
