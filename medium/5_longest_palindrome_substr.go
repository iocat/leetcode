package medium

// DP solution
// Relies on the fact that: cache[i,j] = cache[i+1, j - 1] && charAt[i] == charAt[j] where the cache
// caches whether the substring [i, j] is a palindrome or not.
func longestPalindrome(s string) string {

	switch len(s) {
	case 0:
		return ""
	case 1:
		return s
	}
	var (
		runes          = []rune(s)
		dpIsPalindrome func([]rune, int, int, [][]*bool) bool // checks if the s[i:j]  is a palindrome, dynamically build the cache along the way
		cache          = make([][]*bool, len(runes))
		maxL, maxR     = 0, 0
	)

	for i := range cache {
		cache[i] = make([]*bool, len(runes))
	}

	// recursively checks if the substring runes[i:j+1] is a palindrome
	// also memorize along the way
	dpIsPalindrome = func(runes []rune, i, j int, cache [][]*bool) bool {
		if i == j {
			return true
		} else if i == j-1 {
			return runes[i] == runes[j]
		} else if cache[i][j] != nil && *cache[i][j] {
			return true
		}
		var subPalindrome bool
		if ref := cache[i+1][j-1]; ref == nil {
			subPalindrome = dpIsPalindrome(runes, i+1, j-1, cache)
		} else {
			subPalindrome = *ref
		}
		cache[i][j] = new(bool)
		tmp := subPalindrome && s[i] == s[j]
		*(cache[i][j]) = tmp
		return tmp
	}

	for l := 0; l < len(runes); l++ {
		for r := l; r < len(runes); r++ {
			if dpIsPalindrome(runes, l, r, cache) {
				if r-l > maxR-maxL {
					maxL, maxR = l, r
				}
			}
		}
	}

	return string(runes[maxL : maxR+1])
}

func isAPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	var l, r int
	if length := len(s); length%2 == 0 {
		l, r = (length/2)-1, (length / 2)
	} else {
		l, r = (length-1)/2-1, (length-1)/2+1
	}
	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			return false
		}
		l--
		r++
	}
	return true
}

// Brute Force Solution
func longestPalindrome0(s string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return s
	}

	var max = ""
	for l := 0; l < len(s); l++ {
		for r := l; r < len(s); r++ {
			substr := s[l : r+1]
			if isAPalindrome(substr) && len(substr) > len(max) {
				max = substr
			}
		}
	}
	return max
}
