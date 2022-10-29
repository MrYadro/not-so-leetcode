package leetcode

// https://leetcode.com/problems/ransom-note/
// Runtime: 16 ms
// Memory Usage: 3.9 MB

func CanConstruct(ransomNote string, magazine string) bool {
	lenMag := len(magazine)
	lenRN := len(ransomNote)

	if ransomNote == "" {
		return true
	}

	if lenMag < lenRN {
		return false
	}

	m := make(map[rune]int)

	for i := 0; i < lenMag; i++ {
		m[rune(magazine[i])]++
	}

	for i := 0; i < lenRN; i++ {
		if m[rune(ransomNote[i])] == 0 {
			return false
		}
		m[rune(ransomNote[i])]--
	}

	return true
}
