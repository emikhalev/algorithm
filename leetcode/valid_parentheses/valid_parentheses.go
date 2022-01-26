// https://leetcode.com/problems/valid-parentheses/
package valid_parentheses

func isValid(s string) bool {
	st := NewStack()

	for i := 0; i < len(s); i++ {
		if s[i] == '[' || s[i] == '(' || s[i] == '{' {
			st.Push(s[i])
			continue
		}

		if st.Size() == 0 {
			return false
		}

		pp := st.Pop()
		switch s[i] {
		case ']':
			if pp != '[' {
				return false
			}
		case ')':
			if pp != '(' {
				return false
			}
		case '}':
			if pp != '{' {
				return false
			}
		}
	}
	return st.Size() == 0
}

type Stack struct {
	data []byte
}

func NewStack() *Stack {
	return &Stack{
		data: make([]byte, 0),
	}
}

func (s *Stack) Push(ch byte) {
	s.data = append(s.data, ch)
}

func (s *Stack) Pop() byte {
	v := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return v
}

func (s *Stack) Size() int {
	return len(s.data)
}
