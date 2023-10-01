package main

import (
	"container/heap"
	"fmt"
)

const ResultNum = 30

func main() {
	graph, err := CreateWordsGraph()
	if err != nil {
		fmt.Println(err)
		return
	}

	gem, board, err := GetBoard()
	if err != nil {
		fmt.Println(err)
		return
	}

	h := Solve(board, gem/3, graph)
	result := make([]Word, 0)
	for i := 0; i < h.Len(); i++ {
		if len(result) == 0 {
			result = append(result, heap.Pop(h).(Word))
		} else {
			p := heap.Pop(h).(Word)
			exist := false
			for j := 1; j <= len(result); j++ {
				if p.word == result[len(result)-j].word {
					exist = true
					break
				}
			}

			if !exist {
				result = append(result, p)

				if len(result) >= ResultNum {
					break
				}
			}
		}
	}

	for i := 0; i < len(result); i++ {
		fmt.Printf("%d: \"%s\", %dpoints\n", i+1, result[i].word, result[i].point)
	}
}
