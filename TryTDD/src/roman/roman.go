package roman

func RomanToNunber(a string) int {

	valueI := stringCount(a, "I", 1, []string{"V", "X", "L", "C", "M", "D"})
	valueV := stringCount(a, "V", 5, []string{})
	valueX := stringCount(a, "X", 10, []string{"L", "C", "M", "D"})
	valueL := stringCount(a, "L", 50, []string{})
	valueC := stringCount(a, "C", 100, []string{"M", "D"})
	valueD := stringCount(a, "D", 500, []string{})
	valueM := stringCount(a, "M", 1000, []string{})

	return valueI + valueV + valueX + valueL + valueC + valueD + valueM

}

func stringCount(a string, b string, c int, list []string) int {
	count := 0

	for i, l := 0, len(a); i < l; i++ {
		if string(a[i]) == b {
			d := c
			if i < l-1 {
				if stringInSlice(string(a[i+1]), list) {
					d = d * -1
				}
			}
			count = count + d
		}
	}
	return count
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
