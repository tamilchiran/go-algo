package arraystring

/*
https://leetcode.com/problems/ransom-note?envType=study-plan-v2&envId=top-interview-150
*/
func CanConstruct(ransomNote string, magazine string) bool {
	alpha := [26]int{}
	for _, i := range magazine {
		alpha[i-'a']++
	}

	for _, j := range ransomNote {
		if alpha[j-'a'] <= 0 {
			return false
		}
		alpha[j-'a']--
	}
	return true
}
