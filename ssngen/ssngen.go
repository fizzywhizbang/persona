package ssngen

import (
	"math/rand"
	"strconv"
	"strings"
)

/*
667–679	Not issued
681–699	Not issued
700–728	Railroad Retirement Board (discontinued July 1, 1963)
729–730	Enumeration at entry
*/
//eport validation check
func CheckValid(ssn string) bool {

	firstThree := ssn[0:3]
	//convert first three to int
	num, _ := strconv.Atoi(firstThree)

	if (num >= 666) && (num <= 679) {
		return false
	}
	if (num >= 681) && (num <= 699) {
		return false
	}
	if (num >= 700) && (num <= 730) {
		return false
	}
	//since ssn's before 2011 couldn't be between 734 and 739 or be above 772

	if (num >= 734) && (num <= 739) {
		return false
	}
	if num > 772 {
		return false
	}

	return true
}

//export SSN generator
func SSNgen(qty int) []string {
	//we need to generate three numbers and validate
	var builder strings.Builder
	SSNs := []string{}
	for q := 0; q < qty; q++ {
		for i := 1; i <= 9; i++ {
			if i == 4 || i == 6 {
				builder.WriteString("-")
			}
			d := rand.Intn(9)
			builder.WriteString(strconv.Itoa(d))
		}
		SSNs = append(SSNs, builder.String())
	}

	return SSNs
}
