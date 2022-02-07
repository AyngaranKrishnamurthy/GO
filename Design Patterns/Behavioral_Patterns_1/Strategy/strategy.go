package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

type PrintStrategy interface { //Strategy interface that each printing strategy must implement
	Draw() error
}

type ConsoleSquare struct{}
type ImageSquare struct {
	DestinationFilePath string
}

func (c *ConsoleSquare) Draw() error {
	fmt.Println("Square")
	return nil
}

func (t *ImageSquare) Draw() error {
	//background definition
	width := 800
	height := 600

	origin := image.Point{0, 0}
	bgImage := image.NewRGBA(image.Rectangle{Min: origin, Max: image.Point{X: width, Y: height}})

	bgColor := image.Uniform{color.RGBA{R: 70, G: 70, B: 70, A: 0}}
	quality := &jpeg.Options{Quality: 75}

	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)

	//Square definition
	squareWidth := 200
	squareHeight := 200

	squareColor := image.Uniform{color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{X: (width / 2) - (squareWidth / 2), Y: (height / 2) - (squareHeight / 2)})
	squareImg := image.NewRGBA(square)
	draw.Draw(bgImage, squareImg.Bounds(), &squareColor, origin, draw.Src)

	w, err := os.Create(t.DestinationFilePath)
	if err != nil {
		fmt.Errorf("Error while opening image")
	}
	defer w.Close() //To ensure os.Create is closed after every method call. defer ensures  that the close method is executed even if the error is popped.

	if err = jpeg.Encode(w, bgImage, quality); err != nil {
		fmt.Errorf("Error writing image to disk")
	}

	return nil
}

//A flag is a command the user will pass while using this App. Declaration: "flag.<type>". A flag can have a default value in our case it is the console
var output = flag.String("output", "console", "The output to use between 'console' and 'image' file")

func main() {
	flag.Parse() //While using flags we need to parse it first in the main method as it is important.

	var activeStrategy PrintStrategy

	switch *output {
	case "console":
		activeStrategy = &ConsoleSquare{}
	case "image":
		activeStrategy = &ImageSquare{"/tmp/image.jpg"}
	default:
		activeStrategy = &ConsoleSquare{}
	}

	err := activeStrategy.Draw()
	if err != nil {
		log.Fatal(err)
	}
}

//The ImageSquare is executed and returns no logs to show the execution which is a bad practice. Executed by "go run strategy.go --output image".
//We can't use the above code as a library because it has critical code in the main package hence to overcome the issue  refer strategy_2.
