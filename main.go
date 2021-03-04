package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/anthonynsimon/bild/transform"
)

func main() {
	var (
		fileName      string
		isCorrectFile bool
		printWidth    int
		imgWidth      int
		imgHeight     int
	)

	flag.IntVar(&printWidth, "width", 100, "the number of characters in each row of the printed image")
	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("please provide a jpg or jpeg file to print")
		os.Exit(1)
	}

	if len(flag.Args()) == 0 {
		fileName = os.Args[1]
	} else {
		fileName = flag.Args()[0]
	}

	isCorrectFile = fileName[len(fileName)-3:] == "jpg" || fileName[len(fileName)-4:] == "jpeg"

	if !isCorrectFile {
		fmt.Println("please provide a jpg or jpeg file to print")
		os.Exit(1)
	}

	/////////////////////////////////////////////////////////

	imgFilePath := filepath.Join(fileName)

	img, err := os.Open(imgFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()

	imgData, err := jpeg.Decode(img)
	if err != nil {
		log.Fatal(err)
	}

	imgData = transform.Resize(imgData, printWidth, printWidth*imgData.Bounds().Max.Y/imgData.Bounds().Max.X*45/100, transform.Linear)

	imgWidth = imgData.Bounds().Max.X
	imgHeight = imgData.Bounds().Max.Y

	/////////////////////////////////////////////////////////

	pixels := []string{" ", "░", "▒", "▓", "█"}

	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {
			c := color.GrayModel.Convert(imgData.At(x, y)).(color.Gray)

			pixel := c.Y / 51

			if pixel == 5 {
				pixel--
			}

			fmt.Print(pixels[pixel])
		}

		fmt.Println()
	}
}
