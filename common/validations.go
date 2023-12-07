package common

import (
	"strconv"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func IsNumber(b byte) bool {
	return (b >= 48 && b <= 57)
}

func IsStringNumber(s string) bool {
	result := true
	_, err := strconv.Atoi(s)
	if err != nil {
		result = false
	}
	return result
}
