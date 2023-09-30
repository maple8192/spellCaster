package main

import (
	"bufio"
	"os"
)

type WordsGraph struct {
	initials []*Node
}

type Node struct {
	letter rune
	next   []*Node
}

func CreateWordsGraph() (*WordsGraph, error) {
	root := &WordsGraph{}

	fp, err := os.Open("words_alpha.txt")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		insertWord(root, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if err := fp.Close(); err != nil {
		return nil, err
	}

	return root, nil
}

func insertWord(root *WordsGraph, word []rune) {
	if len(word) <= 1 {
		return
	}

	var node *Node = nil
	for i := 0; i < len(root.initials); i++ {
		if root.initials[i].letter == word[0]-('a'-'A') {
			node = root.initials[i]
			break
		}
	}

	if node == nil {
		node = &Node{word[0] - ('a' - 'A'), make([]*Node, 0)}
		root.initials = append(root.initials, node)
	}

	recInsert(node, word, 1)
}

func recInsert(node *Node, word []rune, p int) {
	if p == len(word) {
		node.next = append(node.next, &Node{' ', make([]*Node, 0)})
		return
	}

	var next *Node = nil
	for i := 0; i < len(node.next); i++ {
		if node.next[i].letter == word[p]-('a'-'A') {
			next = node.next[i]
			break
		}
	}

	if next == nil {
		next = &Node{word[p] - ('a' - 'A'), make([]*Node, 0)}
		node.next = append(node.next, next)
	}

	recInsert(next, word, p+1)
}

func (g *WordsGraph) String() string {
	ret := "WordsGraph ["
	for i := 0; i < len(g.initials); i++ {
		ret += g.initials[i].String()
		if i != len(g.initials)-1 {
			ret += ", "
		}
	}
	return ret + "]"
}

func (n *Node) String() string {
	ret := "Node("
	ret += string(n.letter)
	ret += ") ["
	for i := 0; i < len(n.next); i++ {
		ret += n.next[i].String()
		if i != len(n.next)-1 {
			ret += ", "
		}
	}
	return ret + "]"
}
