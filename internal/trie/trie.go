package trie

import (
	"fmt"
	"sort"
)

type Node struct {
	Words    []string
	Children map[string]*Node
}

func NewNode() *Node {
	n := &Node{Children: map[string]*Node{}}
	return n
}

func (n *Node) AddWords(words []string) {
	for _, word := range words {
		wordSorted := []byte(word)
		sort.Slice(wordSorted, func(i, j int) bool { return wordSorted[i] < wordSorted[j] })
		n.AddWord(word, wordSorted)
	}
}

func (n *Node) AddWord(word string, wordSorted []byte) {
	n.Words = append(n.Words, word)
	if len(wordSorted) == 0 {
		return
	}
	firstLetter := string(wordSorted[0])
	child, ok := n.Children[firstLetter]
	if !ok {
		child = NewNode()
	}
	child.AddWord(word, wordSorted[1:])
	n.Children[firstLetter] = child
}

func (n *Node) Has(has string) []string {
	hasSorted := []byte(has)
	sort.Slice(hasSorted, func(i, j int) bool { return hasSorted[i] < hasSorted[j] })
	fmt.Println("sorted:", string(hasSorted))
	return n.hasHelper(has, hasSorted)
}

func (n *Node) hasHelper(has string, hasSorted []byte) []string {
	fmt.Println(string(hasSorted))
	if len(hasSorted) == 0 {
		return n.allChildWords()
	}
	child, ok := n.Children[string(hasSorted[0])]
	if !ok {
		return []string{}
	}
	return child.hasHelper(has, hasSorted[1:])
}

func (n *Node) allChildWords() []string {
	words := n.Words
	for _, child := range n.Children {
		words = append(words, child.allChildWords()...)
	}
	return words
}
