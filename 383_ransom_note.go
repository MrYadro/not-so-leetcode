package leetcode

// https://leetcode.com/problems/ransom-note/
// Runtime: 8 ms
// Memory Usage: 3.9 MB

func CanConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	m := make([]int, 26)

	for _, c := range magazine {
		m[c-'a']++
	}

	for _, c := range ransomNote {
		adiff := c - 'a'
		if m[adiff] == 0 {
			return false
		}
		m[adiff]--
	}

	return true
}
