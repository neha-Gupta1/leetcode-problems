/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
*/
func isAnagram(s string, t string) bool {
	var original = make(map[rune]int)

	for _, char := range s {
		original[char] = original[char] + 1
	}

	for _, char := range t {
		if _,ok:=original[char] ;ok  {
			original[char] = original[char] - 1
			if original[char] == 0 {
				delete(original, char)
			}
            continue
		}
        if _,ok:=original[char] ;!ok  {
            return false
        }
	}
	if len(original) > 0 {
		return false
	}

	return true
}

/*
time complexity: O(n+m)
space complexity: O(n)
*/