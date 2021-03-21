package util

import (
	"errors"
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

type OutputConfig struct {
	Src          string
	Dst          string
	OutputMode   string
	AsciiPattern string
	OutputWidth  int
	IsUseWeb     bool
	IsPrinted    bool
	IsSaved      bool
	IsQuiet      bool
	IsInverted   bool
}

type ProcessConfig struct {
	Src         string
	IsUseWeb    bool
	OutputWidth int
}

type DrawConfig struct {
	ImgData      image.Image
	ImgWidth     int
	ImgHeight    int
	Src          string
	Dst          string
	OutputMode   string
	AsciiPattern string
	OutputWidth  int
	IsUseWeb     bool
	IsPrinted    bool
	IsSaved      bool
	IsQuiet      bool
	IsInverted   bool
}

func OutputImage(c OutputConfig) (string, error) {
	var (
		src          string = c.Src
		dst          string = c.Dst
		outputMode   string = c.OutputMode
		asciiPattern string = c.AsciiPattern
		outputWidth  int    = c.OutputWidth
		isUseWeb     bool   = c.IsUseWeb
		isPrinted    bool   = c.IsPrinted
		isSaved      bool   = c.IsSaved
		isInverted   bool   = c.IsInverted
		isQuiet      bool   = c.IsQuiet
	)

	var (
		imgData     image.Image
		imgWidth    int
		imgHeight   int
		pixelString string
		err         error
	)

	processOptions := ProcessConfig{
		Src:         src,
		IsUseWeb:    isUseWeb,
		OutputWidth: outputWidth,
	}

	imgData, imgWidth, imgHeight, err = ProcessImage(processOptions)
	if err != nil {
		return "", err
	}

	drawOptions := DrawConfig{
		ImgData:      imgData,
		ImgWidth:     imgWidth,
		ImgHeight:    imgHeight,
		IsSaved:      isSaved,
		Dst:          dst,
		IsInverted:   isInverted,
		OutputMode:   outputMode,
		AsciiPattern: asciiPattern,
		IsPrinted:    isPrinted,
		IsQuiet:      isQuiet,
	}

	pixelString, err = DrawPixels(drawOptions)
	if err != nil {
		return "", err
	}

	return pixelString, nil
}

func ProcessImage(c ProcessConfig) (image.Image, int, int, error) {
	var (
		src         string = c.Src
		isUseWeb    bool   = c.IsUseWeb
		outputWidth int    = c.OutputWidth
	)

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

	imgData = transform.Resize(imgData, outputWidth, outputWidth*imgData.Bounds().Max.Y/imgData.Bounds().Max.X*9/20, transform.Linear)

	imgWidth = imgData.Bounds().Max.X
	imgHeight = imgData.Bounds().Max.Y

	return imgData, imgWidth, imgHeight, nil
}

func GetImgByUrl(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 || res.Header.Get("Content-Type")[:6] != "image/" {
		return nil, errors.New("bad Src provided")
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

func DrawPixels(c DrawConfig) (string, error) {
	var (
		imgData      image.Image = c.ImgData
		imgWidth     int         = c.ImgWidth
		imgHeight    int         = c.ImgHeight
		isSaved      bool        = c.IsSaved
		dst          string      = c.Dst
		isInverted   bool        = c.IsInverted
		outputMode   string      = c.OutputMode
		asciiPattern string      = c.AsciiPattern
		isPrinted    bool        = c.IsPrinted
		isQuiet      bool        = c.IsQuiet
	)

	var (
		pixelLevels string
		pixelLevel  int
		pixelChar   string
		pixelString string
		colored     bool
		progressBar *pb.ProgressBar
		pbTemplate  string
	)

	// 1. have all option logic (ex: cant save in color mode)
	switch outputMode {
	case "ascii":
	case "color":
	case "box":
	default:
		fmt.Println("Please provide a valid print mode. (color, ascii, or box)")
		os.Exit(1)
	}

	if outputMode == "color" {
		if runtime.GOOS == "windows" {
			colored = true
		} else {
			fmt.Println("Color mode not supported.")
			os.Exit(1)
		}
	}
	if outputMode == "box" {
		pixelLevels = " ░▒▓█"
	}
	if outputMode == "ascii" {
		pixelLevels = asciiPattern //  .:-=+*#%@
	}

	if isSaved {
		if colored {
			fmt.Println("Cannot save output in color mode.")
			os.Exit(1)
		}

		if !isPrinted {
			if !isQuiet {
				pbTemplate = `{{ etime . }} {{ bar . "[" "=" ">" " " "]" }} {{speed . }} {{percent . }}`
				progressBar = pb.ProgressBarTemplate(pbTemplate).Start(imgWidth * imgHeight)
			}
		}
	}

	// 2. generate pixelString
	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {
			// applying changes to pixel according to passed params

			l := color.GrayModel.Convert(imgData.At(x, y)).(color.Gray)
			r, g, b, _ := imgData.At(x, y).RGBA()

			if isInverted {
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

			// appending pixel to pixelString or printing pixel

			pixelString += pixelChar

			if isSaved && !isPrinted && !isQuiet {
				progressBar.Increment()
			}
			if isPrinted {
				if colored {
					printColor.RGB(uint8(r), uint8(g), uint8(b), true).Print(pixelChar)
				} else {
					fmt.Print(pixelChar)
				}
			}
		}

		// newline behavior

		pixelString += "\n"

		if isPrinted {
			fmt.Println()
		}
	}

	// 3. handle pixelString according to the passed params (printing to console happens in pixel generation bc it looks cool)

	if isSaved {
		file, err := os.Create(dst)
		if err != nil {
			return "", err
		}
		defer file.Close()

		_, err = file.WriteString(pixelString)
		if err != nil {
			return "", err
		}

		if !isPrinted && !isQuiet {
			progressBar.Finish()
			fmt.Println("Done. Saved to", dst)
		}
	}

	if colored {
		pixelString = ""
	}

	// 4. return pixelString for using DrawPixels outside of imgcli

	return pixelString, nil
}
