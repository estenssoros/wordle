package cmd

import (
	"fmt"
	"sort"

	"github.com/estenssoros/wordle/internal/words"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type Count struct {
	Letter string
	Count  int
}

var commonCmd = &cobra.Command{
	Use:     "common",
	Short:   "outputs a count of letters for all the words in data",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := words.NewClient(data)
		if err != nil {
			return errors.Wrap(err, "words.NewClient")
		}

		counts := make([]Count, len(client.Lookup))
		var i int
		for l, c := range client.Lookup {
			counts[i] = Count{Letter: string(l), Count: c}
			i++
		}
		sort.Slice(counts, func(i, j int) bool { return counts[i].Count > counts[j].Count })
		for _, c := range counts {
			fmt.Println(c.Letter, c.Count)
		}
		return nil
	},
}
