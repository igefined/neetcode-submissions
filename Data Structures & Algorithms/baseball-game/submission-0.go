func calPoints(operations []string) int {
	stack := NewStack()
	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case "+":
			first := stack.Pop()
			second := stack.Pop()
			stack.Push(second)
			stack.Push(first)
			stack.Push(first + second)
		case "C":
			stack.Pop()
		case "D":
			stack.Push(stack.Peek() * 2)
		default:
			parseInt, err := strconv.Atoi(operations[i])
			if err != nil {
				panic(err)
			}
			stack.Push(int32(parseInt))
		}
	}

	var out int32
	for !stack.IsEmpty() {
		out += stack.Pop()
	}
	return int(out)
}

type Stack struct {
	score []int32
}

func (s *Stack) Push(value int32) {
	s.score = append(s.score, value)
}

func (s *Stack) Pop() int32 {
	if len(s.score) == 0 {
		return -1
	}
	lastElem := s.score[len(s.score)-1]
	s.score = s.score[:len(s.score)-1]
	return lastElem
}

func (s *Stack) Peek() int32 {
	if len(s.score) == 0 {
		return -1
	}
	return s.score[len(s.score)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.score) == 0
}

func NewStack() *Stack {
	return &Stack{
		score: make([]int32, 0),
	}
}