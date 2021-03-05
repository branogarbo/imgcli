package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"imgcli/util"
	"io"
	"log"
	"os"

	"github.com/anthonynsimon/bild/transform"
)

func main() {
	var (
		file            string
		isJpg           bool
		isPng           bool
		imgWidth        int
		imgHeight       int
		printWidth      int
		isWebImg        bool
		isAscii         bool
		err             error
		img             io.ReadCloser
		imgData         image.Image
		isPrintSaved    bool
		printSaveTo     string
		isPrintInverted bool
	)

	// process flags/args

	flag.IntVar(&printWidth, "width", 100, "the number of characters in each row of the printed image")
	flag.BoolVar(&isWebImg, "web", false, "whether or not the image is in the filesystem or fetched from the web")
	flag.BoolVar(&isAscii, "ascii", false, "whether or not the the image will be printed as ascii")
	flag.BoolVar(&isPrintSaved, "save", false, "whether or not the the image will be written to a text file")
	flag.BoolVar(&isPrintInverted, "invert", false, "whether or not the the image will be inverted")

	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("please provide a jpg/png file or an image address(url) to print")
		os.Exit(1)
	}

	file = flag.Args()[0]

	if isPrintSaved {
		if len(flag.Args()) == 1 {
			printSaveTo = "print.txt"
		} else {
			printSaveTo = flag.Args()[1]
		}
	}

	if len(file) < 3 {
		fmt.Println("please provide a jpg/png file or an image address(url) to print")
		os.Exit(1)
	}

	isJpg = file[len(file)-3:] == "jpg" || file[len(file)-4:] == "jpeg" || file[len(file)-3:] == "JPG" || file[len(file)-4:] == "JPEG"
	isPng = file[len(file)-3:] == "png" || file[len(file)-3:] == "PNG"

	if !(isJpg || isPng) {
		fmt.Println("please provide a jpg/png file or an image address(url) to print")
		os.Exit(1)
	}

	// process image

	if isWebImg {
		img = util.GetImgByUrl(file)
	} else {
		img = util.GetImgByFilePath(file)
	}
	defer img.Close()

	if isJpg {
		imgData, err = jpeg.Decode(img)
	} else {
		imgData, err = png.Decode(img)
	}

	if err != nil {
		log.Fatal(err)
	}

	imgData = transform.Resize(imgData, printWidth, printWidth*imgData.Bounds().Max.Y/imgData.Bounds().Max.X*45/100, transform.Linear)

	imgWidth = imgData.Bounds().Max.X
	imgHeight = imgData.Bounds().Max.Y

	// draw image

	util.DrawPixels(imgData, imgWidth, imgHeight, isAscii, isPrintSaved, printSaveTo, isPrintInverted)
}
