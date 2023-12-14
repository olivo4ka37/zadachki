package LeetCode

func AsteroidCollision(asteroids []int) []int {
	workedArr := stack1{items: nil}
	xLen := -1

	for i:=0;i<len(asteroids);i++ {
		//fmt.Println(i)
		if asteroids[i] > 0 || len(workedArr.items) == 0 {
			workedArr.push(asteroids[i])
			xLen++
		} else {
			if asteroids[i]*-1 > workedArr.items[xLen] && workedArr.items[xLen] > 0 {
				workedArr.pop()
				xLen--
				i--
			} else if asteroids[i]*-1 < workedArr.items[xLen] && workedArr.items[xLen] > 0 {
				continue
			} else if asteroids[i]*-1 == workedArr.items[xLen] && workedArr.items[xLen] > 0 {
				workedArr.pop()
				xLen--
			} else {
				workedArr.push(asteroids[i])
				xLen++
				//fmt.Println("workedArr.items[xLen] ",workedArr.items[xLen], "asteroids ",asteroids[i])
			}
		}
	}

	asteroids = workedArr.items

	return asteroids
}

type stack1 struct {
	items []int
}

// Push
func (s *stack1) push(i int) {
	s.items = append(s.items, i)
}

// Pop
func (s *stack1) pop() int {
	l := len(s.items) - 1
	lastItem := s.items[l]
	s.items = s.items[:l]
	return lastItem
}
