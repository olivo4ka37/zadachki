package LeetCode

func PredictPartyVictory(senate string) string {
	radiant := false
	dire := false
	ignoreDire := 0
	ignoreRadiant := 0
	result := senate
	if len(senate) == 1 {
		if senate[0] == 82 {
			return "Radiant"
		} else if senate[0] == 68 {
			return "Dire"
		}
	}

	for radiant == false && dire == false {
		radiant = true
		dire = true
		resultNew := ""
		for i := 0; i < len(result); i++ {
			if result[i] == 82 {
				ignoreDire++
				dire = false
				if ignoreRadiant <= 0 {
					resultNew += "R"
				}
				ignoreRadiant--
			} else {
				radiant = false
				ignoreRadiant++
				if ignoreDire <= 0 {
					resultNew += "D"
				}
				ignoreDire--
			}
		}
		result = resultNew
	}

	if radiant {
		return "Radiant"
	} else {
		return "Dire"
	}
}

//Not my O(N) Solution

//func PredictPartyVictory(senate string) string {
//	radiant, dire := make([]int, 0), make([]int, 0)
//
//	for i, party := range(senate) {
//		if party == 'R' {
//			radiant = append(radiant, i)
//		} else {
//			dire = append(dire, i)
//		}
//	}
//
//	for len(radiant) > 0 && len(dire) > 0 {
//		if radiant[0] < dire[0] {
//			radiant = append(radiant, radiant[0] + len(senate))
//		} else {
//			dire = append(dire, dire[0] + len(senate))
//		}
//		radiant = radiant[1:]
//		dire = dire[1:]
//	}
//
//	if len(radiant) > 0 {
//		return "Radiant"
//	} else {
//		return "Dire"
//	}
//}
