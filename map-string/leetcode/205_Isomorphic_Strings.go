package mapString

/*
https://leetcode.com/problems/isomorphic-strings?envType=study-plan-v2&envId=top-interview-150
*/

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapST := make(map[byte]byte)
	mapTS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		ch1 := s[i]
		ch2 := t[i]

		if val, exist := mapST[ch1]; exist {
			if val != ch2 {
				return false
			}
		} else {
			mapST[ch1] = ch2
		}

		if val, exist := mapTS[ch2]; exist {
			if val != ch1 {
				return false
			}
		} else {
			mapTS[ch2] = ch1
		}
	}
	return true
}
