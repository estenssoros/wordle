package cmd

import (
	"fmt"
	"sort"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type Count struct {
	Letter string
	Count  int
}

var commonCmd = &cobra.Command{
	Use:     "common",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		words, err := parseWords()
		if err != nil {
			return errors.Wrap(err, "parseWord")
		}
		lookup := map[string]int{}
		for _, word := range words {
			for _, r := range word {
				lookup[string(r)]++
			}
		}
		counts := make([]Count, len(lookup))
		var i int
		for l, c := range lookup {
			counts[i] = Count{Letter: l, Count: c}
			i++
		}
		sort.Slice(counts, func(i, j int) bool { return counts[i].Count > counts[j].Count })
		for _, c := range counts {
			fmt.Println(c.Letter, c.Count)
		}
		return nil
	},
}
