package util

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/anthonynsimon/bild/transform"
	"github.com/cheggaaa/pb/v3"
	printColor "github.com/gookit/color"
)

func OutputImage(src, dst, outputMode string, outputWidth int, isUseWeb, isSaved, isInverted bool, asciiPattern string) error {
	var (
		imgData   image.Image
		imgWidth  int
		imgHeight int
		err       error
	)

	imgData, imgWidth, imgHeight, err = ProcessImage(src, isUseWeb, outputWidth)
	if err != nil {
		return err
	}

	err = DrawPixels(imgData, imgWidth, imgHeight, isSaved, dst, isInverted, outputMode, asciiPattern)
	if err != nil {
		return err
	}

	return nil
}

func ProcessImage(src string, isUseWeb bool, printWidth int) (image.Image, int, int, error) {
	var (
		img       io.ReadCloser
		imgData   image.Image
		err       error
		imgWidth  int
		imgHeight int
	)

	if isUseWeb {
		img, err = GetImgByUrl(src)
	} else {
		img, err = GetImgByFilePath(src)
	}
	if err != nil {
		return nil, 0, 0, err
	}
	defer img.Close()

	imgData, _, err = image.Decode(img)
	if err != nil {
		return nil, 0, 0, err
	}

	imgData = transform.Resize(imgData, printWidth, printWidth*imgData.Bounds().Max.Y/imgData.Bounds().Max.X*9/20, transform.Linear)

	imgWidth = imgData.Bounds().Max.X
	imgHeight = imgData.Bounds().Max.Y

	return imgData, imgWidth, imgHeight, nil
}

func GetImgByUrl(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}

	return res.Body, nil
}

func GetImgByFilePath(file string) (io.ReadCloser, error) {
	img, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func ScaleValue(value, lowerI, upperI, lowerF, upperF float64) int {
	var (
		initRange     float64
		finalRange    float64
		rangeScale    float64
		relativeValue float64
		scaledValue   float64
	)

	if value > upperI || value < lowerI {
		fmt.Println("Given value is out of the inital range")
		os.Exit(1)
	}

	initRange = upperI - lowerI
	finalRange = upperF - lowerF + 1

	rangeScale = finalRange / initRange
	relativeValue = value - lowerI

	scaledValue = relativeValue*rangeScale + lowerF

	if scaledValue == upperF+1 {
		scaledValue--
	}

	return int(scaledValue)
}

func DrawPixels(imgData image.Image, imgWidth, imgHeight int, isPrintSaved bool, printSaveTo string, isPrintInverted bool, printMode, asciiPattern string) error {
	var (
		pixelLevels     string
		pixelLevel      int
		pixelChar       string
		pixelSaveString string
		colored         bool
		progressBar     *pb.ProgressBar
		pbTemplate      string
	)

	switch printMode {
	case "ascii":
	case "color":
	case "box":
	default:
		fmt.Println("Please provide a valid print mode (color, ascii, or box)")
		os.Exit(1)
	}

	if printMode == "color" {
		if runtime.GOOS == "windows" {
			colored = true
		} else {
			fmt.Println("Color mode not supported.")
			os.Exit(1)
		}
	}
	if printMode == "box" {
		pixelLevels = " ░▒▓█"
	}
	if printMode == "ascii" {
		pixelLevels = asciiPattern //  .:-=+*#%@
	}

	if isPrintSaved {
		pbTemplate = `{{ etime . }} {{ bar . "[" "=" ">" " " "]" }} {{speed . }} {{percent . }}`
		progressBar = pb.ProgressBarTemplate(pbTemplate).Start(imgWidth * imgHeight)
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
				progressBar.Increment()
				pixelSaveString += pixelChar
			} else {
				if colored {
					printColor.RGB(uint8(r), uint8(g), uint8(b), true).Print(pixelChar)
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
			return err
		}
		defer file.Close()

		_, err = file.WriteString(pixelSaveString)
		if err != nil {
			return err
		}

		progressBar.Finish()
		fmt.Println("Done. Saved to", printSaveTo)
	}

	return nil
}

// rewrite DrawPixels to
// 1. have all option logic (ex: cant save in color mode)
// 2. generate pixelString
// 3. handle pixelString according to the passed params (ex: fmt.Print for printing, file.WriteString for saving)
// 4. return pixelString for using DrawPixels outside of imgcli-cobra
