func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var (
		stack = NewStack()
		n     = len(s)
	)

	for i := 0; i < n; i++ {
		el := s[i]
		switch el {
		case '(':
			stack.Push(')')
		case '[':
			stack.Push(']')
		case '{':
			stack.Push('}')
		default:
			if !(stack.Pop() == el) {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

type Stack struct {
	score []uint8
}

func (s *Stack) Push(value uint8) {
	s.score = append(s.score, value)
}

func (s *Stack) Pop() uint8 {
	if len(s.score) == 0 {
		return 0
	}
	lastElem := s.score[len(s.score)-1]
	s.score = s.score[:len(s.score)-1]
	return lastElem
}

func (s *Stack) Peek() uint8 {
	if len(s.score) == 0 {
		return 0
	}
	return s.score[len(s.score)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.score) == 0
}

func NewStack() *Stack {
	return &Stack{
		score: make([]uint8, 0),
	}
}