package main

var points = [...]int{1, 4, 5, 3, 1, 5, 3, 4, 1, 7, 6, 3, 4, 2, 1, 4, 8, 2, 2, 2, 4, 5, 5, 7, 4, 8}

func GetPoint(letter rune) int {
	return points[letter-'A']
}
