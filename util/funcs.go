/*
Copyright © 2021 Brian Longmore brianl.ext@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package util

import (
	"errors"
	"fmt"
	"io"
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
		return nil, errors.New("bad Src provided")
	}

	return res.Body, nil
}

// ScaleValue takes value and transforms(translation and scaling)
// it according to its new, transformed bounds.
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

// isFileExists returns whether filePath already exists.
func isFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	}

	return true
}
