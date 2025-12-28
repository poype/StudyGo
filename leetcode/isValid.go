package main

type Stack struct {
	data []rune
}

func (s *Stack) Push(item rune) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() rune {
	last := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return last
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

type Set struct {
	data map[rune]struct{}
}

func NewSet(items ...rune) *Set {
	set := &Set{data: make(map[rune]struct{}, len(items))}
	for _, v := range items {
		set.data[v] = struct{}{}
	}
	return set
}

func (s *Set) contains(item rune) bool {
	_, ok := s.data[item]
	return ok
}

func isValid(s string) bool {
	stack := Stack{data: make([]rune, 0)}
	openSet := NewSet('(', '[', '{')
	for _, c := range s {
		if openSet.contains(c) {
			stack.Push(c)
		} else {
			if stack.isEmpty() {
				return false
			}
			openC := stack.Pop()
			if openC == '(' && c != ')' {
				return false
			} else if openC == '[' && c != ']' {
				return false
			} else if openC == '{' && c != '}' {
				return false
			}
		}
	}

	return stack.isEmpty()
}
