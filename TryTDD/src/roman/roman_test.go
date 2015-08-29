package roman

import "testing"

func TestRomanToNumber(t *testing.T) {
	testCase := map[string]int{}
	testCase["I"] = 1
	testCase["II"] = 2
	testCase["III"] = 3
	testCase["IV"] = 4
	testCase["V"] = 5
	testCase["VI"] = 6
	testCase["VII"] = 7
	testCase["VIII"] = 8
	testCase["IX"] = 9
	testCase["X"] = 10
	testCase["XI"] = 11
	testCase["XII"] = 12
	testCase["XIII"] = 13
	testCase["XIV"] = 14
	testCase["XV"] = 15
	testCase["XVI"] = 16
	testCase["XVII"] = 17
	testCase["XVIII"] = 18
	testCase["XIX"] = 19
	testCase["XX"] = 20
	testCase["XXX"] = 30
	testCase["XL"] = 40
	testCase["L"] = 50
	testCase["LX"] = 60
	testCase["LXX"] = 70
	testCase["LXXX"] = 80
	testCase["XC"] = 90
	testCase["XCIX"] = 99
	testCase["C"] = 100
	testCase["CI"] = 101
	testCase["CII"] = 102
	testCase["CXCIX"] = 199
	testCase["CC"] = 200
	testCase["CCC"] = 300
	testCase["CD"] = 400
	testCase["D"] = 500
	testCase["DC"] = 600
	testCase["DCCC"] = 800
	testCase["CM"] = 900
	testCase["M"] = 1000
	testCase["MCD"] = 1400
	testCase["MCDXXXVII"] = 1437
	testCase["MD"] = 1500
	testCase["MDCCC"] = 1800
	testCase["MDCCCLXXX"] = 1880
	testCase["MCM"] = 1900
	testCase["MM"] = 2000
	testCase["MMM"] = 3000
	testCase["MMMCCCXXXIII"] = 3333

	for key, value := range testCase {
		r := RomanToNunber(key)
		if r != value {
			t.Errorf("RomanToNunber(%s) failed. Got %d, expected %d.\n", key, r, value)
		}
	}

}
