package rtdata

type Stack struct {
	maxSize int
	size    int
	_top    *Frame
}

func NewStack(maxSize int) *Stack {
	return &Stack{maxSize: maxSize}
}

func (s *Stack) push(frame *Frame) {
	if s.size >= s.maxSize {
		panic("stack over flow error")
	}
	if s._top != nil {
		frame.lower = s._top
	}
	s._top = frame
	s.size++
}

func (s *Stack) pop() *Frame {
	if s._top == nil {
		panic("stack is empty")
	}
	top := s._top
	top.lower = nil
	s._top = s._top.lower
	s.size--
	return top
}

func (s *Stack) top() *Frame {
	if s._top == nil {
		panic("stack is empty")
	}
	return s._top
}
