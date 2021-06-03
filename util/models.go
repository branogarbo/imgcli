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

import "image"

const (
	DefaultMode    = "ascii"
	DefaultPattern = " .,*/(#%&@"
	DefaultWidth   = 100
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
	IsSrcBytes   bool
}

type ProcessConfig struct {
	Src         string
	IsUseWeb    bool
	OutputWidth int
	IsSrcBytes  bool
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
