package main

import (
	"image"
	"image/color"
)


func Name() string {
    return "FilterPlugin"
}

func Apply(img image.Image) image.Image {
    bounds := img.Bounds()
    filteredImg := image.NewRGBA(bounds)

    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            r, g, b, a := img.At(x, y).RGBA()
            filteredImg.SetRGBA(x, y, color.RGBA{R: uint8(g), G: uint8(b), B: uint8(r), A: uint8(a)})
        }
    }

    return filteredImg
}
