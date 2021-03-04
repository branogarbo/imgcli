package main

import (
	"fmt"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/anthonynsimon/bild/transform"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) == 1 {
		log.Fatal("please provide a jpg image to print")
	}

	/////////////////////////////////////////////////////////

	imgFilePath := filepath.Join(cwd, os.Args[1])

	img, err := os.Open(imgFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()

	imgData, err := jpeg.Decode(img)
	if err != nil {
		log.Fatal(err)
	}

	imgData = transform.Resize(imgData, 100, 35, transform.Linear)

	/////////////////////////////////////////////////////////

	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := imgData.Bounds().Min.Y; y < imgData.Bounds().Max.Y; y++ {
		for x := imgData.Bounds().Min.X; x < imgData.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(imgData.At(x, y)).(color.Gray)

			level := c.Y / 51
			if level == 5 {
				level--
			}

			fmt.Print(levels[level])
		}

		fmt.Print("\n")
	}
}
