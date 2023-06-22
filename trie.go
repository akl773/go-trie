package trie

import (
	"errors"
	"fmt"
)

var ErrRecordNotFound = errors.New("record not found")

func formatErrRecordNotFound(word string) error {
	return fmt.Errorf("%s: %w", word, ErrRecordNotFound)
}

type Node struct {
	children map[rune]*Node
	//To store frequency of string in case of duplicates
	count int
}

type Trie struct {
	root *Node
}

func New() *Trie {
	return &Trie{root: &Node{children: make(map[rune]*Node)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &Node{children: make(map[rune]*Node)}
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

func (t *Trie) Delete(word string) error {
	if !t.Search(word) {
		return formatErrRecordNotFound(word)
	}
	t.delete(t.root, word, 0)
	return nil
}

func (t *Trie) delete(node *Node, word string, index int) *Node {
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
