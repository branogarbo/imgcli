package util

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/anthonynsimon/bild/transform"
	"github.com/cheggaaa/pb/v3"
	printColor "github.com/gookit/color"
)

// OutputImage is the high level and preferred method for generating a print.
// It calls both ProcessImage and OutputImage.
func OutputImage(c OutputConfig) (string, error) {
	var (
		src          = c.Src
		dst          = c.Dst
		outputMode   = c.OutputMode
		asciiPattern = c.AsciiPattern
		outputWidth  = c.OutputWidth
		isUseWeb     = c.IsUseWeb
		isPrinted    = c.IsPrinted
		isSaved      = c.IsSaved
		isInverted   = c.IsInverted
		isQuiet      = c.IsQuiet
		isSrcBytes   = c.IsSrcBytes
	)

	var (
		imgData     image.Image
		imgWidth    int
		imgHeight   int
		pixelString string
		err         error
	)

	if outputMode == "" {
		outputMode = DefaultMode
	}
	if outputWidth == 0 {
		outputWidth = DefaultWidth
	}
	if asciiPattern == "" {
		asciiPattern = DefaultPattern
	}

	processOptions := ProcessConfig{
		Src:         src,
		IsUseWeb:    isUseWeb,
		OutputWidth: outputWidth,
		IsSrcBytes:  isSrcBytes,
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

// ProcessImage returns the decoded image and its new dimensions. It is called by OutputImage.
func ProcessImage(c ProcessConfig) (image.Image, int, int, error) {
	var (
		src         = c.Src
		isUseWeb    = c.IsUseWeb
		outputWidth = c.OutputWidth
		isSrcBytes  = c.IsSrcBytes
	)

	var (
		img       io.ReadCloser
		imgData   image.Image
		err       error
		imgWidth  int
		imgHeight int
	)

	if isUseWeb {
		img, err = GetFileByUrl(src)
	} else {
		if isSrcBytes {
			img = ioutil.NopCloser(strings.NewReader(src))
		} else {
			img, err = GetFileByPath(src)
		}
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

// DrawPixels is the low level method for generating a print. It typically uses inputs that would
// be returned by ProcessImage and is called by OutputImage.
func DrawPixels(c DrawConfig) (string, error) {
	var (
		imgData      = c.ImgData
		imgWidth     = c.ImgWidth
		imgHeight    = c.ImgHeight
		isSaved      = c.IsSaved
		dst          = c.Dst
		isInverted   = c.IsInverted
		outputMode   = c.OutputMode
		asciiPattern = c.AsciiPattern
		isPrinted    = c.IsPrinted
		isQuiet      = c.IsQuiet
	)

	var (
		pixelLevels string
		pixelLevel  int
		pixelChar   string
		pixelString string
		colored     bool
		progressBar *pb.ProgressBar
		pbTemplate  string
		err         error
	)

	// 1. have all option logic (ex: cant save in color mode)
	switch outputMode {
	case "ascii":
	case "color":
	case "box":
	default:
		return "", errors.New("please provide a valid print mode. (color, ascii, or box)")
	}

	if outputMode == "color" {
		if runtime.GOOS == "windows" {
			colored = true
		} else {
			return "", errors.New("color mode not supported")
		}
	}
	if outputMode == "box" {
		pixelLevels = " ░▒▓█"
	}
	if outputMode == "ascii" {
		pixelLevels = asciiPattern
	}

	if isSaved {
		if colored {
			return "", errors.New("cannot save output in color mode")
		}

		if !isPrinted {
			if !isQuiet {
				pbTemplate = `{{ etime . }} {{ bar . "[" "=" ">" " " "]" }} {{speed . }} {{percent . }}`
				progressBar = pb.ProgressBarTemplate(pbTemplate).Start(imgWidth * imgHeight).SetMaxWidth(100)
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
					scaledVal, err := ScaleValue(float64(l.Y), 0, 255, 0, float64(len([]rune(pixelLevels))-1))
					if err != nil {
						return "", err
					}

					pixelLevel = len([]rune(pixelLevels)) - 1 - scaledVal
				}
			} else {
				pixelLevel, err = ScaleValue(float64(l.Y), 0, 255, 0, float64(len([]rune(pixelLevels))-1))
				if err != nil {
					return "", err
				}
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
		dst = ProcessFilePath(dst)

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
