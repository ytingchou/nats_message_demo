package cmd

import (
	"github.com/bunyk/gokeybr/app"
	"github.com/bunyk/gokeybr/phrase"
	"github.com/spf13/cobra"
)

var offset, limit int
var textCmd = &cobra.Command{
	Use:     "text [flags] [file with text (\"-\" - stdin)]",
	Aliases: []string{"file"},
	Short:   "train to type contents of some file",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text, skipped, err := phrase.FromFile(args[0], offset, limit)
		fatal(err)

		a, err := app.New(text)
		fatal(err)
		a.Zen = zen
		a.Mute = mute
		a.MinSpeed = minSpeed
		a.Offset = skipped

		a.Run()
		fatal(err)

		saveStats(a, false)

		err = phrase.UpdateFileProgress(args[0], a.LinesTyped(), offset)
		fatal(err)
	},
}

func init() {
	textCmd.Flags().IntVarP(&limit, "length", "l", 0,
		"Minimal lenght in characters of text to train on (default 0 - unlimited)",
	)
	textCmd.Flags().IntVarP(&offset, "offset", "o", -1,
		"Offset in lines when loading file (default 0)",
	)
	rootCmd.AddCommand(textCmd)
}
