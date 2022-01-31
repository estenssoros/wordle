package words

import (
	"github.com/estenssoros/wordle/internal/runes"
)

func Has(words []string, has string) []string {
	if len(has) == 0 {
		return words
	}
	out := []string{}
	for _, word := range words {
		if contains(word, has) {
			out = append(out, word)
		}
	}
	return out
}

func contains(word, has string) bool {
	for _, r := range has {
		if !runes.InWord(word, r) {
			return false
		}
	}
	return true
}

func Excludes(words []string, exclude string) []string {
	if len(exclude) == 0 {
		return words
	}
	out := []string{}
	for _, word := range words {
		if excludes(word, exclude) {
			out = append(out, word)
		}
	}
	return out
}

func excludes(word, exclude string) bool {
	for _, r := range exclude {
		if runes.InWord(word, r) {
			return false
		}
	}
	return true
}

func Orders(words []string, order string) []string {
	if len(order) == 0 {
		return words
	}
	out := []string{}
	for _, word := range words {
		if orders(word, order) {
			out = append(out, word)
		}
	}
	return out
}

func orders(word, order string) bool {
	for i := 0; i < len(word); i++ {
		if !runes.Match(word[i], order[i]) {
			return false
		}
	}
	return true
}

func NotOrders(words, notOrders []string) []string {
	if len(notOrders) == 0 {
		return words
	}
	out := notOrder(words, notOrders[0])
	for _, nO := range notOrders {
		out = notOrder(out, nO)
	}
	return out
}

func notOrder(words []string, nO string) []string {
	if len(nO) == 0 {
		return words
	}
	out := []string{}
	for _, word := range words {
		if doesntMatch(word, nO) {
			out = append(out, word)
		}
	}
	return out
}

func doesntMatch(word, pattern string) bool {
	for i := 0; i < len(word); i++ {
		if runes.DoesntMatch(word[i], pattern[i]) {
			return false
		}
	}
	return true
}
