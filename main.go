package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"

	"github.com/anthonynsimon/bild/transform"
	"github.com/branogarbo/imgcli/util"
)

func main() {
	var (
		file                string
		imgWidth            int
		imgHeight           int
		printWidth          int
		isWebImg            bool
		err                 error
		img                 io.ReadCloser
		imgData             image.Image
		isPrintSaved        bool
		printSaveTo         string
		isPrintInverted     bool
		printMode           string
		asciiPattern        string
		defaultAsciiPattern string = " .-+*#%@"
	)

	// process flags/args

	usageString := `imgcli

A simple tool written in Go that prints images to the command line and more.

Usage:
	imgcli [--mode=<mode>] [--width=<number>] [--invert] [--save] [--ascii] <path-to-image>
	imgcli [--mode=<mode>] [--width=<number>] [--invert] [--save] [--ascii] [--web] "<image-url>"

Flags:`

	flag.IntVar(&printWidth, "width", 100, "the number of characters in each row of the printed image")
	flag.BoolVar(&isWebImg, "web", false, "whether the image is in the filesystem or fetched from the web")
	flag.BoolVar(&isPrintSaved, "save", false, "whether or not the the print will be written to a text file")
	flag.BoolVar(&isPrintInverted, "invert", false, "whether or not the the print will be inverted")
	flag.StringVar(&printMode, "mode", "ascii", "the mode the image will be printed in. (color, ascii, or box)")
	flag.StringVar(&asciiPattern, "ascii", defaultAsciiPattern, "the pattern of ascii characters from least to greatest visibility. pattern of over 8 characters is not recommended")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usageString)
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\n\n", `Use "imgcli --help" or "imgcli -h" to view usage.`)
	}

	flag.Parse()

	switch printMode {
	case "box":
	case "ascii":
	case "color":
	default:
		fmt.Println("please provide a valid print mode (color, ascii, or box)")
		os.Exit(1)
	}

	if asciiPattern != defaultAsciiPattern {
		printMode = "ascii"
	}

	if len(flag.Args()) == 0 {
		fmt.Println("please provide an image file or address(url) to print")
		os.Exit(1)
	}

	file = flag.Args()[0]

	if isPrintSaved {
		if printMode == "color" {
			fmt.Println("cannot save print in color mode.")
			os.Exit(1)
		} else {
			if len(flag.Args()) == 1 {
				printSaveTo = "./print.txt"
			} else {
				printSaveTo = flag.Args()[1]
			}
		}
	}

	if len(file) < 3 {
		fmt.Println("please provide an image file or address(url) to print")
		os.Exit(1)
	}

	// process image

	if isWebImg {
		img = util.GetImgByUrl(file)
	} else {
		img = util.GetImgByFilePath(file)
	}
	defer img.Close()

	imgData, _, err = image.Decode(img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	imgData = transform.Resize(imgData, printWidth, printWidth*imgData.Bounds().Max.Y/imgData.Bounds().Max.X*9/20, transform.Linear)

	imgWidth = imgData.Bounds().Max.X
	imgHeight = imgData.Bounds().Max.Y

	// draw image

	util.DrawPixels(imgData, imgWidth, imgHeight, isPrintSaved, printSaveTo, isPrintInverted, printMode, asciiPattern)
}
