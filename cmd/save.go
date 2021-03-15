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

var saveCmd = &cobra.Command{
	Use:     "save",
	Short:   "Saves output image to a text file.",
	Example: `imgcli save -w 200 -W "https://url-to-some/image.jpg"`,
	Args:    cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		src = args[0]

		if len(args) == 2 {
			dst = args[1]
		} else {
			dst = "./print.txt"
		}

		outputMode, err = cmd.Flags().GetString("mode")
		outputWidth, err = cmd.Flags().GetInt("width")
		isUseWeb, err = cmd.Flags().GetBool("web")
		isInverted, err = cmd.Flags().GetBool("invert")
		asciiPattern, err = cmd.Flags().GetString("ascii")
		isQuiet, err = cmd.Flags().GetBool("quiet")

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
			IsSaved:      true,
			IsInverted:   isInverted,
			AsciiPattern: asciiPattern,
			IsPrinted:    false,
			IsQuiet:      isQuiet,
		}

		_, err = util.OutputImage(options)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
