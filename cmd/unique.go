package cmd

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	uniqueCmd.Flags().BoolVarP(&toClipboard, "clipboard", "", false, "output words to clipboard")
}

var uniqueCmd = &cobra.Command{
	Use:     "unique",
	Short:   "parses data.txt for unique words",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		words, err := parseWords()
		if err != nil {
			return errors.Wrap(err, "parseWords")
		}
		if toClipboard {
			return clipboard.WriteAll(strings.Join(words, "\n"))
		}
		for _, word := range words {
			fmt.Println(word)
		}
		return nil
	},
}
