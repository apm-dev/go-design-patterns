package shapes

import (
	"fmt"
	"github.com/apm-dev/go-design-patterns/behavioral/strategy"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
)

type ImageSquare struct {
	strategy.PrintOutput
}

func (i *ImageSquare) Print() error {
	width := 800
	height := 600

	bgColor := image.Uniform{C: color.RGBA{R: 70, G: 70, B: 70, A: 0}}
	origin := image.Point{X: 0, Y: 0}
	quality := &jpeg.Options{Quality: 75}
	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})
	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)

	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{C: color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImage := image.NewRGBA(square)
	draw.Draw(bgImage, squareImage.Bounds(), &squareColor, origin, draw.Src)

	if i.Writer == nil {
		return fmt.Errorf("no writer stored on ImageSquare")
	}
	if err := jpeg.Encode(i.Writer, squareImage, quality); err != nil {
		return fmt.Errorf("error writing image to disk")
	}
	if i.LogWriter != nil {
		i.LogWriter.Write([]byte("image written in provided writer\n"))
	}
	return nil
}
