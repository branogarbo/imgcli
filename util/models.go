package util

import "image"

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
