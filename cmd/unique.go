package cmd

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/estenssoros/wordle/internal/words"
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
		client, err := words.NewClient(data)
		if err != nil {
			return errors.Wrap(err, "words.NewClient")
		}
		if toClipboard {
			return clipboard.WriteAll(strings.Join(client.Words, "\n"))
		}
		for _, word := range client.Words {
			fmt.Println(word)
		}
		return nil
	},
}
