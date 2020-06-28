// https://leetcode.com/problems/basic-calculator/
package main

import (
	"fmt"
)

func calculate(s string) int {
	s = reverse(s)
	return calc(s)
}


func calc(st string) int {
	bs := []byte(st)
	s := NewStack()
	for i:=0;i<len(bs);i++ {
		// Get number and place in stack
		if isNum(bs[i]) {
			sum := 0
			k := 1
			for ;i<len(bs) && isNum(bs[i]);i++ {
				sum = sum + int(bs[i]-'0')*k
				k *= 10
			}
			s.Push(sum)
			if i>=len(bs)  {
				break
			}
		}
		// operand and closing parenthesis place in stack
		if isOp(bs[i]) || bs[i]==')' {
			s.Push(int(bs[i]))
		}
		// calc
		if bs[i]=='(' {
			sum := s.Pop()

			for ;; {
				op := s.Pop()
				if byte(op) == ')' {
					s.Push(sum)
					break
				}
				b := s.Pop()

				switch byte(op) {
				case  '+':  sum += b
				case  '-':  sum -= b
				}

			}

		}
	}

	//s.Print()
	ans := s.Pop()
	for ;s.Size()>0; {
		op := s.Pop()
		b := s.Pop()
		switch byte(op) {
		case  '+':  ans += b
		case  '-':  ans -= b
		}
	}
	return ans
}

func reverse(s string) string {
	bs := []byte(s)
	l := len(bs)
	for i:=0;i<len(bs)/2;i++ {
		t := bs[i]
		bs[i] = bs[l - 1 - i]
		bs[l - 1 - i] = t
	}
	return string(bs)
}


func isNum(b byte) bool {
	return b>='0' && b<='9'
}

func isOp(b byte) bool {
	return b=='+' || b=='-' || b=='/' || b=='*'
}

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() int {
	idx := len(s.data)-1
	v := s.data[idx]
	s.data = s.data[:idx]
	return v
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Print() {
	for i:=len(s.data)-1;i>=0;i-- {
		v := s.data[i]
		st := string(v)
		if !isOp(byte(v)) {
			st = string(v + '0')
		}
		fmt.Printf("%v", st)
	}
	fmt.Println()
}