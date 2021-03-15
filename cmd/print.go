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
	"fmt"
	"os"

	"github.com/branogarbo/imgcli/util"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:     "print",
	Short:   "Prints output to the command line.",
	Example: "imgcli print --invert ./images/pic.jpg",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		src = args[0]

		outputMode, err = cmd.Flags().GetString("mode")
		outputWidth, err = cmd.Flags().GetInt("width")
		isUseWeb, err = cmd.Flags().GetBool("web")
		isInverted, err = cmd.Flags().GetBool("invert")
		asciiPattern, err = cmd.Flags().GetString("ascii")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		options := util.OutputConfig{
			Src:          src,
			Dst:          dst,
			OutputMode:   outputMode,
			OutputWidth:  outputWidth,
			IsUseWeb:     isUseWeb,
			IsSaved:      false,
			IsInverted:   isInverted,
			AsciiPattern: asciiPattern,
			IsPrinted:    true,
			IsQuiet:      false, // doesnt matter for printing
		}

		_, err = util.OutputImage(options)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
