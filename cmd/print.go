package cmd

import (
	"github.com/branogarbo/imgcli/util"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:     "print",
	Short:   "Prints output to the command line.",
	Example: "imgcli print --invert ./images/pic.jpg",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		src := args[0]

		outputMode, _ := cmd.Flags().GetString("mode")
		outputWidth, _ := cmd.Flags().GetInt("width")
		isUseWeb, _ := cmd.Flags().GetBool("web")
		isInverted, _ := cmd.Flags().GetBool("invert")
		asciiPattern, _ := cmd.Flags().GetString("ascii")

		options := util.OutputConfig{
			Src:          src,
			OutputMode:   outputMode,
			OutputWidth:  outputWidth,
			IsUseWeb:     isUseWeb,
			IsInverted:   isInverted,
			AsciiPattern: asciiPattern,
			IsPrinted:    true,
		}

		_, err := util.OutputImage(options)
		return err
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
