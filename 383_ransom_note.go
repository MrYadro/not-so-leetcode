package leetcode

func CanConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)

	for i := 0; i < len(magazine); i++ {
		m[rune(magazine[i])]++
	}

	for i := 0; i < len(ransomNote); i++ {
		if m[rune(ransomNote[i])] == 0 {
			return false
		}
		m[rune(ransomNote[i])]--
	}

	return true
}
