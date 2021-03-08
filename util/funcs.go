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

func ScaleValue(value, lowerI, upperI, lowerF, upperF float64) int {
	if value > upperI || value < lowerI {
		fmt.Println("given value is out of the inital range")
		os.Exit(1)
	}

	initRange := upperI - lowerI
	finalRange := upperF - lowerF + 1

	rangeScale := finalRange / initRange
	relativeValue := value - lowerI

	scaledValue := relativeValue*rangeScale + lowerF

	if scaledValue == upperF+1 {
		scaledValue--
	}

	return int(scaledValue)
}

func DrawPixels(imgData image.Image, imgWidth, imgHeight int, isPrintSaved bool, printSaveTo string, isPrintInverted bool, printMode, asciiPattern string) {
	var (
		pixelLevels     string
		pixelLevel      int
		pixelChar       string
		pixelSaveString string
		colored         bool
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
		pixelLevels = " ░▒▓█"
	}
	if printMode == "ascii" {
		pixelLevels = asciiPattern //  .:-=+*#%@
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
					pixelLevel = len([]rune(pixelLevels)) - 1 - ScaleValue(float64(l.Y), 0, 255, 0, float64(len([]rune(pixelLevels))-1))
				}
			} else {
				pixelLevel = ScaleValue(float64(l.Y), 0, 255, 0, float64(len([]rune(pixelLevels))-1))
			}

			if colored {
				pixelChar = " "
			} else {
				pixelChar = string([]rune(pixelLevels)[pixelLevel])
			}

			if isPrintSaved {
				pixelSaveString += pixelChar
			} else {
				if colored {
					pcolor.RGB(uint8(r), uint8(g), uint8(b), true).Print(pixelChar)
				} else {
					fmt.Print(pixelChar)
				}
			}
		}

		if isPrintSaved {
			pixelSaveString += "\n"
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

		_, err = file.WriteString(pixelSaveString)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("done. saved to", printSaveTo)
	}
}
