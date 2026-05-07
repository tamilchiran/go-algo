package mapString

/*
https://leetcode.com/problems/ransom-note?envType=study-plan-v2&envId=top-interview-150
*/
func CanConstruct(ransomNote string, magazine string) bool {
	alpha := make(map[rune]int)
	for _, i := range magazine {
		alpha[i]++
	}

	for _, j := range ransomNote {
		alpha[j]--
		if alpha[j] < 0 {
			return false
		}
	}
	return true
}
