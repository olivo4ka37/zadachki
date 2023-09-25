package LeetCode

func Compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	var xMatch byte
	result, schet := 0, 1

	xMatch = chars[0]

	byteNumbers := []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}

	for i := 1; i < len(chars); i++ {
		if xMatch != chars[i] {
			if schet == 1 {
				chars[result] = xMatch
				result++
				xMatch = chars[i]
				schet = 0
			} else if schet < 10 && schet > 1 {
				chars[result] = xMatch
				chars[result+1] = byteNumbers[schet]
				result += 2
				xMatch = chars[i]
				schet = 0
			} else if schet >= 10 && schet < 100 {
				chars[result] = xMatch
				chars[result+1] = byteNumbers[(schet / 10)]
				chars[result+2] = byteNumbers[(schet % 10)]
				result += 3
				xMatch = chars[i]
				schet = 0
			} else if schet >= 100 && schet < 1000 {
				chars[result] = xMatch
				chars[result+1] = byteNumbers[(schet / 100)]
				chars[result+2] = byteNumbers[((schet / 100) / 10)]
				chars[result+3] = byteNumbers[(schet % 10)]
				result += 4
				xMatch = chars[i]
				schet = 0
			} else if schet >= 1000 && schet < 10000 {
				chars[result] = xMatch
				chars[result+1] = byteNumbers[(schet / 1000)]
				chars[result+2] = byteNumbers[((schet / 1000) / 100)]
				chars[result+3] = byteNumbers[(((schet / 1000) / 100) / 10)]
				chars[result+4] = byteNumbers[(schet % 10)]
				result += 5
				xMatch = chars[i]
				schet = 0
			}
		}
		schet++
	}

	if schet == 1 {
		chars[result] = chars[len(chars)-1]
		result++
	} else if schet < 10 && schet > 1 {
		chars[result] = xMatch
		chars[result+1] = byteNumbers[schet]
		result += 2
	} else if schet >= 10 && schet < 100 {
		chars[result] = xMatch
		chars[result+1] = byteNumbers[(schet / 10)]
		chars[result+2] = byteNumbers[(schet % 10)]
		result += 3
	} else if schet >= 100 && schet < 1000 {
		chars[result] = xMatch
		chars[result+1] = byteNumbers[(schet / 100)]
		chars[result+2] = byteNumbers[((schet / 100) / 10)]
		chars[result+3] = byteNumbers[(schet % 10)]
		result += 4
	} else if schet >= 1000 && schet < 10000 {
		chars[result] = xMatch
		chars[result+1] = byteNumbers[(schet / 1000)]
		chars[result+2] = byteNumbers[((schet / 1000) / 100)]
		chars[result+3] = byteNumbers[(((schet / 1000) / 100) / 10)]
		chars[result+4] = byteNumbers[(schet % 10)]
		result += 5
	}

	return result
}
