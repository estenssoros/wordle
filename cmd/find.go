package cmd

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/estenssoros/wordle/internal/runes"
	"github.com/estenssoros/wordle/internal/words"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	has       string
	order     string
	exclude   string
	notOrders []string
)

func init() {
	findCmd.Flags().StringVarP(&has, "has", "", "", "letters in word")
	findCmd.Flags().StringVarP(&order, "order", "", "", "order of the word (? is unknown)")
	findCmd.Flags().StringVarP(&exclude, "exclude", "", "", "letters to exclude")
	findCmd.Flags().StringArrayVarP(&notOrders, "not-order", "", []string{}, "letters in the word but in wrong order (? is unknown)")
}

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(has) > 5 {
			return errors.New("has is greater than 5")
		}
		if len(order) > 5 {
			return errors.New("order is greater than 5")
		}
		for _, notOrder := range notOrders {
			if len(notOrder) != 5 {
				return errors.Errorf("--not-order: %s does not have length 5", notOrder)
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		options, err := parseWords()
		if err != nil {
			return errors.Wrap(err, "parseWords")
		}
		if len(has) == 0 && len(order) == 0 && len(exclude) == 0 && len(notOrders) == 0 {
			fmt.Println("random word:", randomWord(options))
			return nil
		}
		has = addOrder(has, order)
		choices := words.NotOrders(words.Orders(words.Excludes(words.Has(options, has), exclude), order), notOrders)
		if len(choices) == 0 {
			return errors.New("no words found")
		}
		fmt.Println(strings.Repeat("-", 50))
		if len(choices) == 1 {
			fmt.Println(choices[0])
			fmt.Println(strings.Repeat("-", 50))
			return nil
		}
		fmt.Printf("try %s\n", randomWord(choices))
		fmt.Println(strings.Repeat("-", 50))
		for _, choice := range choices {
			fmt.Println(choice)
		}
		fmt.Println(strings.Repeat("-", 50))
		return nil
	},
}

func addOrder(has, order string) string {
	for _, r := range order {
		if r == '?' {
			continue
		}
		if !runes.InWord(has, r) {
			has += string(r)
		}
	}
	return has
}

func randomWord(words []string) string {
	rand.Seed(time.Now().Unix())
	return words[rand.Intn(len(words))]
}
