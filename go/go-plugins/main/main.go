package main

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"plugin"
)


func main() {
    // Load the plugin
    plug, err := plugin.Open(filepath.Join(".", "myfilter.so"))
    if err != nil {
        panic(err)
    }

    // Lookup the symbol that implements the FilterPlugin interface
    sym, err := plug.Lookup("Apply")
    if err != nil {
        panic(err)
    }

    // Assert that the symbol is of the correct type
    filterPlugin, ok := sym.(func(img image.Image) image.Image)
    if !ok {
        panic("Plugin does not implement the FilterPlugin interface")
    }

    // Apply the filter to an image
    img := image.NewRGBA(image.Rect(0, 0, 100, 100))
    filteredImg := filterPlugin(img)

    // Save the filtered image
    outputFile, err := os.Create("output.png")
    if err != nil {
        panic(err)
    }
    defer outputFile.Close()

    png.Encode(outputFile, filteredImg)
}
