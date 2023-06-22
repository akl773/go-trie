# Trie Package in Go

This is a Trie data structure implemented in the Go programming language. A Trie, also known as a prefix tree, is a type of search tree that is used to store a dynamic set or associative array where the keys are usually strings. 

This Trie package allows you to create a Trie, add words, search for words, and delete words.

## Getting Started

To use this Trie package, you need to have Go installed on your machine. For instructions on how to install Go, refer to the [official Go documentation](https://golang.org/doc/install).

## Usage

Here is an example of how to use the Trie package:

```go
package main

import (
	"fmt"
 	"github.com/akl773/go-trie/trie"
)

func main() {
	t := trie.New()
	t.Insert("go")
	t.Insert("golang")
	t.Insert("gopher")
	t.Insert("godoc")
	t.Insert("good")

	fmt.Println(t.Search("go"))       // True
	fmt.Println(t.Search("gopher"))   // True
	fmt.Println(t.Search("godzilla")) // False

	err := t.Delete("gopher")
	if err!=nil {
		if errors.Is(err, ErrRecordNotFound) {
			// invalid request
        }
		// handle 
    }
	fmt.Println(t.Search("gopher")) // False
}
```

## API
`NewTrie()`: Create a new Trie.

`Insert(word string)`: Insert a word into the Trie.

`Search(word string)`: Check if a word is in the Trie.

`Delete(word string) error`: Delete a word from the Trie.


## Error Handling
If a record is not found, the package will return a ErrRecordNotFound error. This can be checked using the standard errors.Is function from the Go library.
