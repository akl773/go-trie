package main

import (
	"errors"
	"fmt"
)

var ErrRecordNotFound = errors.New("record not found")

func FormatErrRecordNotFound(word string) error {
	return fmt.Errorf("%s: %w", word, ErrRecordNotFound)
}

type TrieNode struct {
	children map[rune]*TrieNode
	//To store frequency of string in case of duplicates
	count int
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.count++
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return node.count != 0
}

func (t *Trie) Delete(word string) {
	t.delete(t.root, word, 0)
}

func (t *Trie) delete(node *TrieNode, word string, index int) *TrieNode {
	if index == len(word) {
		if node.count != 0 {
			node.count--
		}
		if len(node.children) == 0 {
			node = nil
		}
		return node
	}

	ch := rune(word[index])
	if child, ok := node.children[ch]; ok {
		node.children[ch] = t.delete(child, word, index+1)
	}

	if len(node.children) == 0 && node.count == 0 {
		node = nil
	}

	return node
}

func main() {
	t := NewTrie()
	t.Insert("go")
	t.Insert("golang")
	t.Insert("gopher")
	t.Insert("godoc")
	t.Insert("good")

	fmt.Println(t.Search("go"))
	fmt.Println(t.Search("gopher"))
	fmt.Println(t.Search("godzilla"))

	t.Delete("gopher")
	fmt.Println(t.Search("gopher"))
}
