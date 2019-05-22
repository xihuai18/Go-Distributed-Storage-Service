package trie

import (
	"fmt"
)

type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{children: make(map[byte]*TrieNode), isEnd: false}
}

type Trie struct{
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: newTrieNode()}
}

func (trie *Trie) Insert(word string) {
	node := trie.root
	for i := 0; i < len(word); i++ {
		_, ok := node.children[word[i]]
		if !ok {
			node.children[word[i]] = newTrieNode()
		}
		node = node.children[word[i]]
	}
	node.isEnd = true
}

func (trie *Trie) Search(word string) bool {
	node := trie.root
	for i := 0; i < len(word); i++ {
		_, ok := node.children[word[i]]
		if !ok {
			return false
		}
		node = node.children[word[i]]
	}
	return node.isEnd
}

func (trie *Trie) StartsWith(prefix string) *TrieNode {
	node := trie.root
	for i := 0; i < len(prefix); i++ {
		_, ok := node.children[prefix[i]]
		if !ok {
			return nil
		}
		node = node.children[prefix[i]]
	}
	return node
}

func (trie *Trie) IsStartWith(prefix string) bool {
	node := trie.root
	for i := 0; i < len(prefix); i++ {
		_, ok := node.children[prefix[i]]
		if !ok {
			return false
		}
		node = node.children[prefix[i]]
	}
	return true
}

func (trie *Trie) AllStartWith(prefix string) []string {
	ret := make([]string, 0)
	node := trie.StartsWith(prefix)
	if node == nil {
		return make([]string, 0)
	}
	node2string := make(map[*TrieNode]string)
	node2string[node] = fmt.Sprintf("%v", prefix)
	que := make([]*TrieNode, 0)
	que = append(que, node)
	for len(que) > 0 {
		node = que[len(que)-1]
		que = que[:len(que)-1]
		if node.isEnd == true {
			ret = append(ret, node2string[node])
		}
		for alp, nd := range node.children {
			que = append(que, nd)
			node2string[nd] = node2string[node]+fmt.Sprintf("%s", string(alp))
		}
	}
	return ret
}
