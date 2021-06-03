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
package cmd

import (
	u "github.com/branogarbo/imgcli/util"
	"github.com/spf13/cobra"
)

var (
	isUseWeb     bool
	isInverted   bool
	outputMode   string
	outputWidth  int
	asciiPattern string
	src          string
	dst          string
	err          error
	isQuiet      bool
)

var rootCmd = &cobra.Command{
	Use:   "imgcli",
	Short: "A simple tool written in Go that prints images to the command line and more.",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	AddOutputFlags(printCmd)
	AddOutputFlags(saveCmd)

	saveCmd.Flags().BoolVarP(&isQuiet, "quiet", "q", false, "Whether or not the save output is quiet")
}

func AddOutputFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&isUseWeb, "web", "W", false, "Whether the source image is in the filesystem or fetched from the web")
	cmd.Flags().BoolVarP(&isInverted, "invert", "i", false, "Whether or not the the print will be inverted")
	cmd.Flags().StringVarP(&outputMode, "mode", "m", u.DefaultMode, "he mode the image will be printed in")
	cmd.Flags().IntVarP(&outputWidth, "width", "w", u.DefaultWidth, "The number of characters in each row of the output")
	cmd.Flags().StringVarP(&asciiPattern, "ascii", "p", u.DefaultPattern, "The pattern of ascii characters from least to greatest visibility")
}
