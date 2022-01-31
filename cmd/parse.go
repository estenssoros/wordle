package cmd

import (
	"bufio"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

type Node struct {
	Children map[byte]*Node
}

func parseWords() ([]string, error) {
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
