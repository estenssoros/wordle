package trie

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAddWord(t *testing.T) {
	words := []string{
		"sugar",
		"perky",
		"alter",
		"altos",
		"alums",
	}
	n := NewNode()
	n.AddWords(words)
	ju, err := json.MarshalIndent(n, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ju))
	options := n.Has("alt")
	fmt.Println(options)
}
