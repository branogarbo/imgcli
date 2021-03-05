package util

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"net/http"
	"os"
)

func GetImgByUrl(url string) io.ReadCloser {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		log.Fatal(err)
	}

	return res.Body
}

func GetImgByFilePath(file string) io.ReadCloser {
	img, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func DrawPixels(imgData image.Image, imgWidth, imgHeight int, isAscii bool) {
	var pixels []string

	if isAscii {
		pixels = []string{" ", ".", ",", "~", "#"}
	} else {
		pixels = []string{" ", "░", "▒", "▓", "█"}
	}

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
