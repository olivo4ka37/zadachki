package LeetCode

func IsSubsequence(s string, t string) bool {
	
	if len(s) == 0 {
		return true
	}

	xNumberOfS := 0
	for i:=0;i<len(t);i++ {
		if xNumberOfS == len(s) {
			break
		} else if s[xNumberOfS] == t[i] {
			xNumberOfS++
		}
	}
	
	if xNumberOfS == len(s) {
		return true
	} else {
		return false
	}
}
