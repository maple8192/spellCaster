package main

type LetterStack []*Letter

func (s *LetterStack) Push(x *Letter) {
	*s = append(*s, x)
}

func (s *LetterStack) Pop() *Letter {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *LetterStack) Last() *Letter {
	return (*s)[len(*s)-1]
}

func (s *LetterStack) Exist(x *Letter) bool {
	for i := 0; i < len(*s); i++ {
		if (*s)[i].letter == x.letter && (*s)[i].lm == x.lm && (*s)[i].wm == x.wm {
			return true
		}
	}

	return false
}

func (s *LetterStack) String() string {
	ret := "["
	for i := 0; i < len(*s); i++ {
		ret += string((*s)[i].letter)
		if i != len(*s)-1 {
			ret += ", "
		}
	}
	return ret + "]"
}
