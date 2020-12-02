package lib

// ProcessOne handles logic for part one of question
func ProcessOne() int {
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

//8-11 x: xxrbhxskxxf
type passwordCheck struct {
	password string
	char     byte
	min      int
	max      int
}
