package cmd

import (
	"errors"

	"github.com/branogarbo/imgcli/util"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:     "save",
	Short:   "Saves output image to a text file.",
	Example: `imgcli save -w 200 -W https://url-to-some/image.jpg`,
	Args:    cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("image source not provided")
		}

		src := args[0]
		var dst string

		if len(args) == 2 {
			dst = args[1]
		} else {
			dst = "./print.txt"
		}

		outputMode, _ := cmd.Flags().GetString("mode")
		outputWidth, _ := cmd.Flags().GetInt("width")
		isUseWeb, _ := cmd.Flags().GetBool("web")
		isInverted, _ := cmd.Flags().GetBool("invert")
		asciiPattern, _ := cmd.Flags().GetString("ascii")
		isQuiet, _ := cmd.Flags().GetBool("quiet")

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

		_, err := util.OutputImage(options)
		return err
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
