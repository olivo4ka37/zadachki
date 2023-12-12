package LeetCode

func RemoveStars(s string) string {
	workedStr := stack{items: nil}

	for i := 0; i < len(s); i++ {
		if string(s[i]) == `*` {
			workedStr.pop()
		} else {
			workedStr.push(s[i])
		}
	}

	s = string(workedStr.items)

	return s
}

type stack struct {
	items []byte
}

// Push
func (s *stack) push(i byte) {
	s.items = append(s.items, i)
}

// Pop
func (s *stack) pop() byte {
	l := len(s.items) - 1
	lastItem := s.items[l]
	s.items = s.items[:l]
	return lastItem
}
