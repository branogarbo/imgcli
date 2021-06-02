package cmd

import (
	"fmt"

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
			return
		}

		options := util.OutputConfig{
			Src:          src,
			OutputMode:   outputMode,
			OutputWidth:  outputWidth,
			IsUseWeb:     isUseWeb,
			IsInverted:   isInverted,
			AsciiPattern: asciiPattern,
			IsPrinted:    true,
		}

		_, err = util.OutputImage(options)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
