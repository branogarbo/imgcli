/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"os"

	"github.com/branogarbo/imgcli-cobra/util"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:     "print",
	Short:   "Prints images to the command line",
	Example: "imgcli print -w 200 ./images/pic.jpg",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		src = args[0]

		isUseWeb, err = cmd.Flags().GetBool("web")
		outputWidth, err = cmd.Flags().GetInt("width")
		isInverted, err = cmd.Flags().GetBool("invert")
		outputMode, err = cmd.Flags().GetString("mode")
		asciiPattern, err = cmd.Flags().GetString("ascii")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		imgData, imgWidth, imgHeight, err = util.ProcessImage(src, isUseWeb, outputWidth)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		util.DrawPixels(imgData, imgWidth, imgHeight, false, "", isInverted, outputMode, asciiPattern)
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
