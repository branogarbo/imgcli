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

	"github.com/branogarbo/imgcli-cobra/util"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:     "save",
	Short:   "Saves converted image to a text file",
	Example: "imgcli-cobra save -i ./images/pic.jpg",
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

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		_, err = util.OutputImage(src, dst, outputMode, outputWidth, isUseWeb, true, isInverted, asciiPattern, false)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
