package Trie

import "fmt"

func main() {
	myTrie := initTrie()

	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, v := range toAdd {
		myTrie.insert(v)
	}

	fmt.Println(myTrie.search("hui"))
}

// alphabetSize is the number of possible characters in the trie
const alphabetSize = 26

// node represents each node in the trie
type node struct {
	children [alphabetSize]*node
	isEnd    bool
}

// Trie represents the trie and has the pointer
type trie struct {
	root *node
}

// initTrie will create new Trie
func initTrie() *trie {
	result := &trie{root: &node{}}
	return result
}

// insert will take in a word and add it to the trie
func (t *trie) insert(word string) {
	wordLength := len(word)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// search will take in a word and RETURN true is that word is included in the trie
func (t *trie) search(word string) bool {
	wordLength := len(word)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}

		currentNode = currentNode.children[charIndex]
	}

	return currentNode.isEnd
}
