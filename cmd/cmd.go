package cmd

import (
	_ "embed"

	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(findCmd)
	cmd.AddCommand(crawlCmd)
	cmd.AddCommand(uniqueCmd)
	cmd.AddCommand(commonCmd)
}

//go:embed data.txt
var data string

var cmd = &cobra.Command{
	Use:   "wordle",
	Short: "tries to solve the daily wordle!",
}

func Execute() error {
	return cmd.Execute()
}
