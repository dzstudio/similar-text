package similartxt

// return the len of longest string both in str1 and str2 and the positions in str1 and str2
func SimilarStr(str1 []rune, str2 []rune) (int, int, int) {
	var maxLen, tmp, pos1, pos2 = 0, 0, 0, 0
	len1, len2 := len(str1), len(str2)

	for p := 0; p < len1; p++ {
		for q := 0; q < len2; q++ {
			tmp = 0
			for p+tmp < len1 && q+tmp < len2 && str1[p+tmp] == str2[q+tmp] {
				tmp++
			}
			if tmp > maxLen {
				maxLen, pos1, pos2 = tmp, p, q
			}
		}

	}

	return maxLen, pos1, pos2
}

// return the total length of longest string both in str1 and str2
func SimilarChar(str1 []rune, str2 []rune) int {
	maxLen, pos1, pos2 := SimilarStr(str1, str2)
	total := maxLen

	if maxLen != 0 {
		if pos1 > 0 && pos2 > 0 {
			total += SimilarChar(str1[:pos1], str2[:pos2])
		}
		if pos1+maxLen < len(str1) && pos2+maxLen < len(str2) {
			total += SimilarChar(str1[pos1+maxLen:], str2[pos2+maxLen:])
		}
	}

	return total
}

// return a int value in [0, 100], which stands for match level
func SimilarText(str1 string, str2 string) int {
	txt1, txt2 := []rune(str1), []rune(str2)
	if len(txt1) == 0 || len(txt2) == 0 {
		return 0
	}
	return SimilarChar(txt1, txt2) * 200 / (len(txt1) + len(txt2))
}
