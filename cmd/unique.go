package cmd

import (
	"strings"

	"github.com/atotto/clipboard"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var uniqueCmd = &cobra.Command{
	Use:     "unique",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		words, err := parseWords()
		if err != nil {
			return errors.Wrap(err, "parseWords")
		}
		clipboard.WriteAll(strings.Join(words, "\n"))
		return nil
	},
}
