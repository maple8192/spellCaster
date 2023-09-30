package main

type IntStack []int

func (s *IntStack) Push(x int) {
	*s = append(*s, x)
}

func (s *IntStack) Pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *IntStack) Last() int {
	return (*s)[len(*s)-1]
}

func (s *IntStack) Exist(x int) bool {
	for i := 0; i < len(*s); i++ {
		if (*s)[i] == x {
			return true
		}
	}

	return false
}
