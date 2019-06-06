package trie

type TrieNode struct {
	Children map[rune]*TrieNode
	IsWord   bool
}

type Trie struct {
	Root TrieNode
}

func (t *Trie) Insert(s string) {
	current := &t.Root
	for _, r := range s {
		if _, ok := current.Children[r]; !ok {
			current.Children = make(map[rune]*TrieNode)
			current.Children[r] = &TrieNode{make(map[rune]*TrieNode), false}
		}
		current = current.Children[r]
	}

	current.IsWord = true
}

func (t *Trie) Exists(s string) bool {
	current := &t.Root
	for _, r := range s {
		if _, ok := current.Children[r]; !ok {
			return false
		}
		current = current.Children[r]
	}

	return current.IsWord
}
