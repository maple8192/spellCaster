package main

import (
	"container/heap"
	"fmt"
)

func Solve(board *[25]Letter, wildcards int, words *WordsGraph) *WordHeap {
	hp := &WordHeap{}

	Foreach(25, wildcards, func(indices *[]int) {
		if len(*indices) != 0 {
			f := true
			for i := 0; i < len(*indices)-1; i++ {
				if (*indices)[i] != (*indices)[i+1]-1 {
					f = false
					break
				}
			}
			if f {
				fmt.Printf("[ %d / %d ] ...\n", (*indices)[0], 25-len(*indices))
			}
		}

		for i := 0; i < 25; i++ {
			for j := 0; j < len(words.initials); j++ {
				exist := false
				for k := 0; k < len(*indices); k++ {
					if (*indices)[k] == i {
						exist = true
						break
					}
				}

				if exist {
					search(board, words.initials[j], indices, &IntStack{i}, &LetterStack{&Letter{words.initials[j].letter, board[i].lm, board[i].wm}}, hp)
				} else {
					if words.initials[j].letter == board[i].letter {
						search(board, words.initials[j], indices, &IntStack{i}, &LetterStack{&board[i]}, hp)
						break
					}
				}
			}
		}
	})

	return hp
}

func search(board *[25]Letter, node *Node, wildcards *[]int, path *IntStack, letters *LetterStack, hp *WordHeap) {
	for i := 0; i < len(node.next); i++ {
		if node.next[i].letter == ' ' {
			runes := make([]rune, 0)
			points := 0
			multiply := 1
			for j := 0; j < len(*letters); j++ {
				runes = append(runes, (*letters)[j].letter)
				points += GetPoint((*letters)[j].letter) * (*letters)[j].lm
				multiply *= (*letters)[j].wm
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
				if node.next[j].letter != ' ' {
					path.Push(next[i])
					letters.Push(&Letter{node.next[j].letter, board[next[i]].lm, board[next[i]].wm})
					search(board, node.next[j], wildcards, path, letters, hp)
					letters.Pop()
					path.Pop()
				}
			} else {
				if node.next[j].letter == board[next[i]].letter {
					path.Push(next[i])
					letters.Push(&board[next[i]])
					search(board, node.next[j], wildcards, path, letters, hp)
					letters.Pop()
					path.Pop()
					break
				}
			}
		}
	}
}
