package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Letter struct {
	letter rune
	lm     int
	wm     int
}

var ErrCannotParse = errors.New("parse error")

func GetBoard() (int, *[25]Letter, error) {
	var in string
	if _, err := fmt.Scan(&in); err != nil {
		return 0, nil, err
	}
	gem, err := strconv.Atoi(in)
	if err != nil {
		return 0, nil, err
	}

	if _, err := fmt.Scan(&in); err != nil {
		return 0, nil, err
	}

	var letters []Letter

	for i := 0; i < len(in); i++ {
		if 'A' <= in[i] && in[i] <= 'Z' {
			letters = append(letters, Letter{rune(in[i]), 1, 1})
			continue
		} else if 'a' <= in[i] && in[i] <= 'z' {
			letters = append(letters, Letter{rune(in[i]) - ('a' - 'A'), 1, 1})
			continue
		} else if in[i] == '!' {
			if len(letters) != 0 {
				letters[len(letters)-1].lm = 2
				continue
			}
		} else if in[i] == '@' {
			if len(letters) != 0 {
				letters[len(letters)-1].lm = 3
				continue
			}
		} else if in[i] == '#' {
			if len(letters) != 0 {
				letters[len(letters)-1].wm = 2
				continue
			}
		} else if in[i] == '$' {
			if len(letters) != 0 {
				letters[len(letters)-1].wm = 3
				continue
			}
		}

		return 0, nil, ErrCannotParse
	}

	if len(letters) != 25 {
		return 0, nil, ErrCannotParse
	}

	var board [25]Letter
	copy(board[:], letters)

	return gem, &board, nil
}
