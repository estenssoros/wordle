package words

import (
	"bufio"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/estenssoros/wordle/internal/runes"
	"github.com/pkg/errors"
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

// LetterCounts counts letters in each word
func LetterCounts(words []string) map[rune]int {
	lookup := map[rune]int{}
	for _, word := range words {
		for _, r := range word {
			lookup[r]++
		}
	}
	return lookup
}

// Client words client
type Client struct {
	Words  []string
	Lookup map[rune]int
}

// NewClient creats a new client
func NewClient(data string) (*Client, error) {
	words, err := parseWords(data)
	if err != nil {
		return nil, errors.Wrap(err, "parseWords")
	}
	return &Client{
		Words:  words,
		Lookup: LetterCounts(words),
	}, nil
}

// Has update client words with has
func (c *Client) Has(has string) *Client {
	c.Words = Has(c.Words, has)
	return c
}

// Excludes update client words with excludes
func (c *Client) Excludes(excludes string) *Client {
	c.Words = Excludes(c.Words, excludes)
	return c
}

// Orders update client words with orders
func (c *Client) Orders(orders string) *Client {
	c.Words = Orders(c.Words, orders)
	return c
}

// NotOrders update client words with no orders
func (c *Client) NotOrders(notOrders []string) *Client {
	c.Words = NotOrders(c.Words, notOrders)
	return c
}

// Choice a choice option
type Choice struct {
	Word  string
	Score int
}

// Rank ranks words according to their letter occurrence
func (c *Client) Rank() *Client {
	choices := make([]Choice, len(c.Words))
	for i, word := range c.Words {
		choices[i] = Choice{Word: word, Score: c.rankWord(word)}
	}
	sort.Slice(choices, func(i, j int) bool { return choices[i].Score > choices[j].Score })
	for i, choice := range choices {
		c.Words[i] = choice.Word
	}
	return c
}

// rankWord rank a word by it's letters
func (c *Client) rankWord(word string) int {
	var ttl int
	for _, r := range word {
		ttl += c.Lookup[r]
	}
	return ttl
}

// RandomWord choose a random word
func (c *Client) RandomWord() string {
	rand.Seed(time.Now().Unix())
	return c.Words[rand.Intn(len(c.Words))]
}

// parseWords parses data into words
func parseWords(data string) ([]string, error) {
	unique := map[string]struct{}{}
	s := bufio.NewScanner(strings.NewReader(data))
	for s.Scan() {
		word := s.Text()
		if len(word) != 5 {
			return nil, errors.Errorf("%s is not length 5", word)
		}
		unique[strings.ToLower(word)] = struct{}{}
	}
	words := make([]string, len(unique))
	var i int
	for word := range unique {
		words[i] = word
		i++
	}
	sort.Strings(words)
	return words, nil
}
