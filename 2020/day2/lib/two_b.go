package lib

// ProcessTwo handles logic for part two of question
func ProcessTwo() int {
	passwords := importFile("pw.txt")
	var valid, cnt int
top:
	for _, vals := range passwords {
		cnt = 0
		for _, char := range vals.password {
			if char == rune(vals.char) {
				cnt++
			}
			if cnt > vals.max {
				continue top
			}
		}
		if cnt >= vals.min {
			valid++
		}
	}

	return valid
}

