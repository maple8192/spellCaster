package main

func Foreach(n, k int, f func(*[]int)) {
	indices := make([]int, k)
	recForeach(&indices, n-1, k, f)
}

func recForeach(indices *[]int, s, r int, f func(*[]int)) {
	if r == 0 {
		f(indices)
	} else {
		if s < 0 {
			return
		}
		recForeach(indices, s-1, r, f)
		(*indices)[r-1] = s
		recForeach(indices, s-1, r-1, f)
	}
}
