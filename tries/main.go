package main

import (
	"fmt"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

func newTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func (t *Trie) Insert(word string) {
	currentNode := t.root

	for _, char := range word {
		_, ok := currentNode.children[char]
		if !ok {
			currentNode.children[char] = NewTrieNode()
		}
		currentNode = currentNode.children[char]
	}

	currentNode.isEnd = true
}

func (t *Trie) Search(word string) bool {
	currentNode := t.root
	for _, char := range word {
		if _, ok := currentNode.children[char]; !ok {
			return false
		}
		currentNode = currentNode.children[char]
	}

	return currentNode.isEnd
}

func main() {
	trie := newTrie()
	words := []string{"apple", "app", "banana", "band"}

	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println(trie.Search("banan"))
}
