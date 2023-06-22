package trie

import (
	"errors"
	"testing"
)

func TestTrie_InsertSearchDelete(t *testing.T) {
	t.Run("empty trie", func(t *testing.T) {
		trie := New()
		if trie.Search("test") {
			t.Errorf("Expected empty trie to return false for any search, got true")
		}
	})

	t.Run("insert and search", func(t *testing.T) {
		trie := New()
		words := []string{"test", "tester", "testing"}

		for _, word := range words {
			trie.Insert(word)
			if !trie.Search(word) {
				t.Errorf("Expected trie to contain inserted word %q", word)
			}
		}
	})

	t.Run("delete", func(t *testing.T) {
		trie := New()
		words := []string{"test", "tester", "testing"}

		for _, word := range words {
			trie.Insert(word)
		}

		trie.Delete("testing")
		if trie.Search("testing") {
			t.Errorf("Expected trie to not contain deleted word, got true")
		}

		trie.Delete("test")
		if trie.Search("test") {
			t.Errorf("Expected trie to not contain deleted word, got true")
		}
	})

	t.Run("delete non-existing word", func(t *testing.T) {
		trie := New()
		err := trie.Delete("nonExisting")
		if !errors.Is(err, ErrRecordNotFound) {
			t.Errorf("Expected ErrRecordNotFound error, got different error or nil")
		}
	})
}
