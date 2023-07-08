package CodeWars

func ArrayDiff(a, b []int) []int {
	boolMap := make([]bool, len(a))
	var result []int

	for i := 0; i < len(a); i++ {
		boolMap[i] = true
	}

	for i := 0; i < len(a); i++ {
		for c := 0; c < len(b); c++ {
			if a[i] == b[c] {
				boolMap[i] = false
			}
		}
	}

	for i := 0; i < len(a); i++ {
		if boolMap[i] == true {
			result = append(result, a[i])
		}
	}

	return result
}
