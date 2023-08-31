package LeetCode

import "fmt"

func Compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	var xMatch byte
	result, schet := 0, 0

	xMatch = chars[0]

	for i := 1; i < len(chars); i++ {
		fmt.Println("result:", result)
		fmt.Println("chars:", string(chars[result:]))
		schet++
		if xMatch != chars[i] {
			if schet == 1 {
				schet = 0
				chars[result] = xMatch
				result++
			} else if schet < 10 && schet > 1 {
				chars[result+1] = xMatch
				chars[result+2] = byte(schet)
				result += 2
			} else if schet >= 10 && schet < 100 {
				chars[result+1] = xMatch
				chars[result+2] = byte(schet / 10)
				chars[result+3] = byte(schet % 10)
				result += 3
			} else if schet >= 100 && schet < 1000 {
				chars[result+1] = xMatch
				chars[result+2] = byte(schet / 100)
				chars[result+2] = byte((schet / 100) / 10)
				chars[result+3] = byte(schet % 10)
				result += 4
			} else if schet >= 1000 && schet < 10000 {
				chars[result+1] = xMatch
				chars[result+2] = byte(schet / 1000)
				chars[result+2] = byte((schet / 1000) / 100)
				chars[result+2] = byte(((schet / 1000) / 100) / 10)
				chars[result+3] = byte(schet % 10)
				result += 5
			}
		}
	}

	//fmt.Println(string(chars[:result]))
	return result
}
