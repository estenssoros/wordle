package cmd

import (
	"fmt"
	"strings"

	"github.com/estenssoros/wordle/internal/runes"
	"github.com/estenssoros/wordle/internal/words"
	"github.com/fatih/color"
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
	Short: "attempts to find the solution to the daily wordle",
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
		client, err := words.NewClient(data)
		if len(has) == 0 && len(order) == 0 && len(exclude) == 0 && len(notOrders) == 0 {
			fmt.Printf("random word: ")
			color.Green(client.RandomWord())
			fmt.Println()
			return nil
		}
		has = addNotOrder(addOrder(has, order), notOrders)
		if err != nil {
			return errors.Wrap(err, "words.NewClient")
		}
		client = client.Has(has).Excludes(exclude).Orders(order).NotOrders(notOrders).Rank()
		if len(client.Words) == 0 {
			return errors.New("no words found")
		}

		fmt.Println(strings.Repeat("-", 50))
		defer fmt.Println(strings.Repeat("-", 50))

		if len(client.Words) == 1 {
			color.Green(client.Words[0])
			return nil
		}
		for _, choice := range client.Words {
			color.Green(choice)
		}
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

func addNotOrder(has string, notOrders []string) string {
	for _, notOrder := range notOrders {
		has = addOrder(has, notOrder)
	}
	return has
}
