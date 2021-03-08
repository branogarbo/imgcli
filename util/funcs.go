package util

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"net/http"
	"os"
	"runtime"

	pcolor "github.com/gookit/color"
)

func GetImgByUrl(url string) io.ReadCloser {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		fmt.Println(err)
		os.Exit(1)
	}

	return res.Body
}

func GetImgByFilePath(file string) io.ReadCloser {
	img, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return img
}

func DrawPixels(imgData image.Image, imgWidth, imgHeight int, isPrintSaved bool, printSaveTo string, isPrintInverted bool, printMode string) {
	var (
		pixelLevel  []string
		pixelString string
		pixel       uint8
		colored     bool
	)

	if printMode == "color" {
		if runtime.GOOS == "windows" {
			colored = true
		} else {
			fmt.Println("colors not supported.")
			os.Exit(1)
		}
	}
	if printMode == "box" {
		pixelLevel = []string{" ", "░", "▒", "▓", "█"}
	}
	if printMode == "ascii" {
		pixelLevel = []string{" ", ".", ",", "~", "#"}
	}

	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {
			l := color.GrayModel.Convert(imgData.At(x, y)).(color.Gray)
			r, g, b, _ := imgData.At(x, y).RGBA()

			if isPrintInverted {
				if colored {
					r = 255 - r
					g = 255 - g
					b = 255 - b
				} else {
					pixel = (255 - l.Y) / 51 // gives different vals when simplified to 5 - c.Y/51
				}
			} else {
				pixel = l.Y / 51
			}

			if pixel == 5 {
				pixel--
			}

			if isPrintSaved {
				pixelString += pixelLevel[pixel]
			} else {
				if colored {
					pcolor.RGB(uint8(r), uint8(g), uint8(b), true).Print(" ")
				} else {
					fmt.Print(pixelLevel[pixel])
				}
			}
		}

		if isPrintSaved {
			pixelString += "\n"
		} else {
			fmt.Println()
		}
	}

	if isPrintSaved {
		file, err := os.Create(printSaveTo)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		_, err = file.WriteString(pixelString)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("done. saved to", printSaveTo)
	}
}
