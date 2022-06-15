package util

import (
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path"
)

// GetFileByPath returns the file contents from filePath.
// Make sure you close the file after calling.
func GetFileByPath(filePath string) (io.ReadCloser, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// GetFileByUrl returns the http response body from url.
// Make sure you close the body after calling.
func GetFileByUrl(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		msg := fmt.Sprintf("bad Src provided, image fetch returned status code of: %d", res.StatusCode)

		return nil, errors.New(msg)
	}

	return res.Body, nil
}

// ScaleValue takes value and transforms(translation and scaling)
// it according to its new, transformed bounds.
func ScaleValue(value, lowerI, upperI, lowerF, upperF float64) (int, error) {
	var (
		initRange     float64
		finalRange    float64
		rangeScale    float64
		relativeValue float64
		scaledValue   float64
	)

	if value > upperI || value < lowerI {
		return 0, errors.New("given value is out of the inital range")
	}

	initRange = upperI - lowerI
	finalRange = upperF - lowerF

	relativeValue = value - lowerI
	rangeScale = finalRange / initRange

	scaledValue = rangeScale*relativeValue + lowerF

	return int(math.Round(scaledValue)), nil
}

// ProcessFilePath returns the path of a unique destination file in the
// directory of filePath. If filePath does not already exist, it will just
// return filePath. Otherwise it would return a unique path in the
// form of "n_fileName" where n is a positive integer.
func ProcessFilePath(filePath string) string {
	var (
		fileDir, fileName = path.Split(filePath)
		newFileName       = fileName
		newFilePath       = filePath
	)

	for rc := 1; isFileExists(newFilePath); rc++ {
		newFileName = fmt.Sprintf("%v_%v", rc, fileName)

		newFilePath = path.Join(fileDir, newFileName)
	}

	return newFilePath
}

// isFileExists returns whether or not filePath already exists.
func isFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	}

	return true
}
