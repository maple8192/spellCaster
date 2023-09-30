package main

import (
	"container/heap"
)

func Solve(board *[25]Letter, wildcards int, words *WordsGraph) *WordHeap {
	hp := &WordHeap{}

	Foreach(25, wildcards, func(indices *[]int) {
		for i := 0; i < 25; i++ {
			for j := 0; j < len(words.initials); j++ {
				if words.initials[j].letter == board[i].letter {
					search(board, words.initials[j], indices, &IntStack{i}, hp)
					break
				}
			}
		}
	})

	return hp
}

func search(board *[25]Letter, node *Node, wildcards *[]int, path *IntStack, hp *WordHeap) {
	for i := 0; i < len(node.next); i++ {
		if node.next[i].letter == ' ' {
			runes := make([]rune, 0)
			points := 0
			multiply := 1
			for j := 0; j < len(*path); j++ {
				runes = append(runes, board[(*path)[j]].letter)
				points += GetPoint(board[(*path)[j]].letter) * board[(*path)[j]].lm
				multiply *= board[(*path)[j]].wm
			}
			long := 0
			if len(runes) >= 6 {
				long = 10
			}
			heap.Push(hp, Word{string(runes), points*multiply + long})
			break
		}
	}

	next := make([]int, 0)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if 0 <= path.Last()%5+j && path.Last()%5+j < 5 && 0 <= path.Last()/5+i && path.Last()/5+i < 5 {
				if !path.Exist(path.Last() + (5 * i) + j) {
					next = append(next, path.Last()+(5*i)+j)
				}
			}
		}
	}

	for i := 0; i < len(next); i++ {
		for j := 0; j < len(node.next); j++ {
			exist := false
			for k := 0; k < len(*wildcards); k++ {
				if (*wildcards)[k] == next[i] {
					exist = true
					break
				}
			}

			if exist {
				if node.next[j].letter == ' ' {
					path.Push(next[i])
					search(board, node.next[j], wildcards, path, hp)
					path.Pop()
				}
			} else {
				if node.next[j].letter == board[next[i]].letter {
					path.Push(next[i])
					search(board, node.next[j], wildcards, path, hp)
					path.Pop()
					break
				}
			}
		}
	}
}
