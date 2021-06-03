package cmd

import (
	"fmt"

	"github.com/branogarbo/imgcli/util"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:     "save",
	Short:   "Saves output image to a text file.",
	Example: `imgcli save -w 200 -W "https://url-to-some/image.jpg"`,
	Args:    cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please provide an image source")
			return
		}

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
			return
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
			IsQuiet:      isQuiet,
		}

		_, err = util.OutputImage(options)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
